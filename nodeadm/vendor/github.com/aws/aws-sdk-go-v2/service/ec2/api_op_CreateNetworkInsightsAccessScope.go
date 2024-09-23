// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a Network Access Scope.
//
// Amazon Web Services Network Access Analyzer enables cloud networking and cloud
// operations teams to verify that their networks on Amazon Web Services conform to
// their network security and governance objectives. For more information, see the [Amazon Web Services Network Access Analyzer Guide]
// .
//
// [Amazon Web Services Network Access Analyzer Guide]: https://docs.aws.amazon.com/vpc/latest/network-access-analyzer/
func (c *Client) CreateNetworkInsightsAccessScope(ctx context.Context, params *CreateNetworkInsightsAccessScopeInput, optFns ...func(*Options)) (*CreateNetworkInsightsAccessScopeOutput, error) {
	if params == nil {
		params = &CreateNetworkInsightsAccessScopeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateNetworkInsightsAccessScope", params, optFns, c.addOperationCreateNetworkInsightsAccessScopeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateNetworkInsightsAccessScopeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateNetworkInsightsAccessScopeInput struct {

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. For more information, see [How to ensure idempotency].
	//
	// [How to ensure idempotency]: https://docs.aws.amazon.com/ec2/latest/devguide/ec2-api-idempotency.html
	//
	// This member is required.
	ClientToken *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The paths to exclude.
	ExcludePaths []types.AccessScopePathRequest

	// The paths to match.
	MatchPaths []types.AccessScopePathRequest

	// The tags to apply.
	TagSpecifications []types.TagSpecification

	noSmithyDocumentSerde
}

type CreateNetworkInsightsAccessScopeOutput struct {

	// The Network Access Scope.
	NetworkInsightsAccessScope *types.NetworkInsightsAccessScope

	// The Network Access Scope content.
	NetworkInsightsAccessScopeContent *types.NetworkInsightsAccessScopeContent

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateNetworkInsightsAccessScopeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateNetworkInsightsAccessScope{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateNetworkInsightsAccessScope{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateNetworkInsightsAccessScope"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addSpanRetryLoop(stack, options); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addIdempotencyToken_opCreateNetworkInsightsAccessScopeMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateNetworkInsightsAccessScopeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateNetworkInsightsAccessScope(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = addSpanInitializeStart(stack); err != nil {
		return err
	}
	if err = addSpanInitializeEnd(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestStart(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestEnd(stack); err != nil {
		return err
	}
	return nil
}

type idempotencyToken_initializeOpCreateNetworkInsightsAccessScope struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateNetworkInsightsAccessScope) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateNetworkInsightsAccessScope) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateNetworkInsightsAccessScopeInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateNetworkInsightsAccessScopeInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateNetworkInsightsAccessScopeMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateNetworkInsightsAccessScope{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateNetworkInsightsAccessScope(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateNetworkInsightsAccessScope",
	}
}
