package provider

import (
	iam "github.com/pulumi/pulumi-google-native/sdk/go/google/iam/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The set of arguments for creating a StaticPage component resource.
type WorkloadIdentityPoolForGithubArgs struct {
	IdentityPoolName pulumi.StringInput `pulumi:"identityPoolName"`
}

// The StaticPage component resource.
type WorkloadIdentityPoolForGithub struct {
	pulumi.ResourceState

	IdentityPool *iam.WorkloadIdentityPool `pulumi:"identityPool"`
}

// NewStaticPage creates a new StaticPage component resource.
func NewGoogleWorkloadIdentityPoolForGithub(ctx *pulumi.Context,
	name string, args *WorkloadIdentityPoolForGithubArgs, opts ...pulumi.ResourceOption) (*WorkloadIdentityPoolForGithub, error) {
	if args == nil {
		args = &WorkloadIdentityPoolForGithubArgs{}
	}

	component := &WorkloadIdentityPoolForGithub{}
	err := ctx.RegisterComponentResource("github-credentials:google:WorkloadIdentityPoolForGithub", name, component, opts...)
	if err != nil {
		return nil, err
	}

	// either args.IdentityPoolName or "GitHub" should be used
	identityPoolName := args.IdentityPoolName
	if identityPoolName == nil {
		identityPoolName = pulumi.String("GitHub")
	}

	identityPool, err := iam.NewWorkloadIdentityPool(ctx, name, &iam.WorkloadIdentityPoolArgs{
		Disabled:               pulumi.Bool(false),
		DisplayName:            pulumi.String("Github"),
		Description:            pulumi.String("Github Workload Identity Pool"),
		WorkloadIdentityPoolId: identityPoolName,
	}, pulumi.Parent(component))

	if err != nil {
		return nil, err
	}

	component.IdentityPool = identityPool

	if err := ctx.RegisterResourceOutputs(component, pulumi.Map{
		"identityPool": identityPool,
	}); err != nil {
		return nil, err
	}

	return component, nil
}
