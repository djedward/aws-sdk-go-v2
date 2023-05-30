// Code generated by smithy-go-codegen DO NOT EDIT.

package securitylake

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/securitylake/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Automatically enables Amazon Security Lake for new member accounts in your
// organization. Security Lake is not automatically enabled for any existing member
// accounts in your organization.
func (c *Client) CreateDataLakeOrganizationConfiguration(ctx context.Context, params *CreateDataLakeOrganizationConfigurationInput, optFns ...func(*Options)) (*CreateDataLakeOrganizationConfigurationOutput, error) {
	if params == nil {
		params = &CreateDataLakeOrganizationConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDataLakeOrganizationConfiguration", params, optFns, c.addOperationCreateDataLakeOrganizationConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDataLakeOrganizationConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDataLakeOrganizationConfigurationInput struct {

	// Enable Security Lake with the specified configuration settings, to begin
	// collecting security data for new accounts in your organization.
	//
	// This member is required.
	AutoEnableNewAccount []types.DataLakeAutoEnableNewAccountConfiguration

	noSmithyDocumentSerde
}

type CreateDataLakeOrganizationConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDataLakeOrganizationConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateDataLakeOrganizationConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDataLakeOrganizationConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateDataLakeOrganizationConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDataLakeOrganizationConfiguration(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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
	return nil
}

func newServiceMetadataMiddleware_opCreateDataLakeOrganizationConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securitylake",
		OperationName: "CreateDataLakeOrganizationConfiguration",
	}
}