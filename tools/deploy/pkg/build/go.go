package build

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/shadracnicholas/home-automation/libraries/go/exe"
	"github.com/shadracnicholas/home-automation/libraries/go/oops"
	"github.com/shadracnicholas/home-automation/tools/deploy/pkg/config"
	"github.com/shadracnicholas/home-automation/tools/deploy/pkg/git"
	"github.com/shadracnicholas/home-automation/tools/deploy/pkg/output"
	"github.com/shadracnicholas/home-automation/tools/libraries/env"
)

// GoBuilder is a builder for golang
type GoBuilder struct {
	Service *config.Service
	Target  *config.Target
}

// Build build a go binary for the target architecture and puts it in workingDir
func (b *GoBuilder) Build(revision, workingDir string) (*Release, error) {
	if err := git.Init(revision); err != nil {
		return nil, oops.WithMessage(err, "failed to initialise git mirror")
	}

	op := output.Info("Parsing service's config")
	runtimeEnv, err := env.Parse(b.Service.EnvFiles...)
	if err != nil {
		op.Failed()
		return nil, oops.WithMessage(err, "failed to parse service's env files")
	}
	op.Complete()

	op = output.Info("Compiling binary for %s", b.Target.Architecture)

	// Make sure the service exists in the mirror
	pkgToBuild := fmt.Sprintf("./%s", b.Service.Name)
	if _, err := os.Stat(filepath.Join(git.Dir(), pkgToBuild)); err != nil {
		op.Failed()
		return nil, oops.WithMessage(err, "failed to stat service directory")
	}

	binName := b.Service.DashedName()

	buildEnv := os.Environ()
	switch b.Target.Architecture {
	case config.ArchARMv6:
		buildEnv = append(buildEnv, "GOOS=linux", "GOARCH=arm", "GOARM=6")
		binName += "-armv6"
	default:
		op.Failed()
		return nil, oops.InternalService("unsupported architecture %q", b.Target.Architecture)
	}

	hash, err := git.CurrentHash(false)
	if err != nil {
		op.Failed()
		return nil, oops.WithMessage(err, "failed to get hash")
	}

	shortHash, err := git.CurrentHash(true)
	if err != nil {
		op.Failed()
		return nil, oops.WithMessage(err, "failed to get short hash")
	}

	binName = fmt.Sprintf("%s-%s", binName, shortHash)
	binOut := filepath.Join(workingDir, binName)

	flags := fmt.Sprintf("-X github.com/shadracnicholas/home-automation/libraries/go/router.Revision=%s", hash)

	if err := exe.Command("go", "build", "-o", binOut, "-ldflags", flags, pkgToBuild).
		Dir(git.Dir()).Env(buildEnv).Run().Err; err != nil {
		op.Failed()
		return nil, oops.WithMessage(err, "failed to compile")
	}

	op.Complete()

	return &Release{
		Cmd:       binName,
		Env:       runtimeEnv,
		Revision:  hash,
		ShortHash: shortHash,
	}, nil
}
