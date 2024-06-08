// Code generated by smithy-go-codegen DO NOT EDIT.

package ecrpublic

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Notifies Amazon ECR that you intend to upload an image layer.
//
// When an image is pushed, the InitiateLayerUpload API is called once for each
// image layer that hasn't already been uploaded. Whether an image layer uploads is
// determined by the BatchCheckLayerAvailability API action.
//
// This operation is used by the Amazon ECR proxy and is not generally used by
// customers for pulling and pushing images. In most cases, you should use the
// docker CLI to pull, tag, and push images.
func (c *Client) InitiateLayerUpload(ctx context.Context, params *InitiateLayerUploadInput, optFns ...func(*Options)) (*InitiateLayerUploadOutput, error) {
	if params == nil {
		params = &InitiateLayerUploadInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "InitiateLayerUpload", params, optFns, c.addOperationInitiateLayerUploadMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*InitiateLayerUploadOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type InitiateLayerUploadInput struct {

	// The name of the repository that you want to upload layers to.
	//
	// This member is required.
	RepositoryName *string

	// The Amazon Web Services account ID, or registry alias, that's associated with
	// the registry to which you intend to upload layers. If you do not specify a
	// registry, the default public registry is assumed.
	RegistryId *string

	noSmithyDocumentSerde
}

type InitiateLayerUploadOutput struct {

	// The size, in bytes, that Amazon ECR expects future layer part uploads to be.
	PartSize *int64

	// The upload ID for the layer upload. This parameter is passed to further UploadLayerPart and CompleteLayerUpload
	// operations.
	UploadId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationInitiateLayerUploadMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpInitiateLayerUpload{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpInitiateLayerUpload{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "InitiateLayerUpload"); err != nil {
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
	if err = addOpInitiateLayerUploadValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opInitiateLayerUpload(options.Region), middleware.Before); err != nil {
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
	return nil
}

func newServiceMetadataMiddleware_opInitiateLayerUpload(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "InitiateLayerUpload",
	}
}
