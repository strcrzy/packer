package s3upload

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/s3"
	"github.com/mitchellh/packer/common"
	"github.com/mitchellh/packer/packer"
)

const BuilderId = "packer.post-processor.s3upload"

type Config struct {
	common.PackerConfig `mapstructure:",squash"`

	AccessKey  string `mapstructure:"access_key"`
	SecretKey  string `mapstructure:"secret_key"`
	BucketName string `mapstructure:"bucket"`
	Region     string
	Path       string

	tpl *packer.ConfigTemplate
}

type PostProcessor struct {
	config Config
}

func (p *PostProcessor) Configure(raws ...interface{}) error {
	_, err := common.DecodeConfig(&p.config, raws...)
	if err != nil {
		return err
	}

	p.config.tpl, err = packer.NewConfigTemplate()
	if err != nil {
		return err
	}
	p.config.tpl.UserVars = p.config.PackerUserVars

	// Accumulate any errors
	errs := new(packer.MultiError)

	templates := map[string]*string{
		"access_key":  &p.config.AccessKey,
		"secret_key":  &p.config.SecretKey,
		"region":      &p.config.Region,
		"bucket_name": &p.config.BucketName,
		"path":        &p.config.Path,
	}

	for key, ptr := range templates {
		if *ptr == "" {
			errs = packer.MultiErrorAppend(
				errs, fmt.Errorf("%s must be set", key))
		}

		*ptr, err = p.config.tpl.Process(*ptr, nil)
		if err != nil {
			errs = packer.MultiErrorAppend(
				errs, fmt.Errorf("Error processing %s: %s", key, err))
		}
	}

	_, err := aws.GetAuth(p.config.AccessKey, p.config.SecretKey)
	if err != nil {
		errs = packer.MultiErrorAppend(
			errs, fmt.Errorf("Error processing AWS credentials: %s", err))
	}
	
	if len(errs.Errors) > 0 {
		return errs
	}

	return nil

}

func (p *PostProcessor) PostProcess(ui packer.Ui, artifact packer.Artifact) (packer.Artifact, bool, error) {
	auth, err := aws.GetAuth(p.config.AccessKey, p.config.SecretKey)
	if err != nil {
		return nil, false, err
	}
	region := aws.Regions[p.config.Region]
	conn := s3.New(auth, region)
	bucket := conn.Bucket(p.config.BucketName)
	for _, path := range artifact.Files() {
		ui.Message("Uploading artifact: " + path)
		contents, err := ioutil.ReadFile(path)
		if err != nil {
			return nil, false, err
		}
		s3path := p.config.Path + "/" + path
		err = bucket.Put(s3path, contents, "application/x-tar", s3.Private)
		if err != nil {
			return nil, false, err
		}

		fileUrl := bucket.URL(s3path)

		ui.Message("Uploaded artifact to s3: " + fileUrl)

	 }

	// Build the artifact
	artifact = S3UploadArtifact{
		S3BucketValue:  p.config.BucketName,
		S3PathValue:    p.config.Path,
		FilesValue:			artifact.Files(),
	}

	return artifact, false, nil
}
