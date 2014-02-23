package s3upload

import (
  "bytes"
  "fmt"
)

// S3UploadArtifact is an Artifact implementation for when an image is uploaded
// to an Amazon S3 Bucket.
type S3UploadArtifact struct {
  IdValue        string
  BuilderIdValue string
  S3BucketValue  string
  S3PathValue    string
  FilesValue     []string
}

func (a *S3UploadArtifact) Id() string {
  return a.IdValue
}

func (a *Artifact) BuilderId() string {
  return BuilderId
}

func (a *S3UploadArtifact) BuilderId() string {
  return a.BuilderIdValue
}

func (a *S3UploadArtifact) Files() []string {
  return a.FilesValue
}

func (a *S3UploadArtifact) String() string {
  var buf bytes.Buffer
  for i, str := range a.FilesValue {
    if(i != 0) {
      buf.WriteString(", ")
    }
    buf.WriteString(str)
  }
  return fmt.Sprintf("Uploaded files: %s.", buf)
}

func (a *S3UploadArtifact) Destroy() error {
  return nil
}
