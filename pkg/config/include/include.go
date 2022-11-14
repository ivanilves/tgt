package include

import (
	"os"
	"strings"
)

func fileOrSymlink(d os.DirEntry) bool {
	return d.Type().IsRegular() || d.Type() == os.ModeSymlink
}

// IsTerragrunt tells us if we operate on Terragrunt config file
func IsTerragrunt(d os.DirEntry) bool {
	return fileOrSymlink(d) && d.Name() == "terragrunt.hcl"
}

// IsTerraform tells us if we operate on Terraform file(s)
func IsTerraform(d os.DirEntry) bool {
	return fileOrSymlink(d) && strings.HasSuffix(d.Name(), ".tf")
}

// IsTerraformOrTerragrunt tells us if we operate on Terraform or Terragrunt file(s)
func IsTerraformOrTerragrunt(d os.DirEntry) bool {
	return IsTerraform(d) || IsTerragrunt(d)
}
