// *** WARNING: this file was generated by the Lumi IDL Compiler (LUMIDL). ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
    "errors"

    pbempty "github.com/golang/protobuf/ptypes/empty"
    pbstruct "github.com/golang/protobuf/ptypes/struct"
    "golang.org/x/net/context"

    "github.com/pulumi/lumi/pkg/resource"
    "github.com/pulumi/lumi/pkg/resource/plugin"
    "github.com/pulumi/lumi/pkg/tokens"
    "github.com/pulumi/lumi/pkg/util/contract"
    "github.com/pulumi/lumi/pkg/util/mapper"
    "github.com/pulumi/lumi/sdk/go/pkg/lumirpc"
)

/* RPC stubs for SecurityGroup resource provider */

// SecurityGroupToken is the type token corresponding to the SecurityGroup package type.
const SecurityGroupToken = tokens.Type("aws:ec2/securityGroup:SecurityGroup")

// SecurityGroupProviderOps is a pluggable interface for SecurityGroup-related management functionality.
type SecurityGroupProviderOps interface {
    Check(ctx context.Context, obj *SecurityGroup) ([]error, error)
    Create(ctx context.Context, obj *SecurityGroup) (resource.ID, error)
    Get(ctx context.Context, id resource.ID) (*SecurityGroup, error)
    InspectChange(ctx context.Context,
        id resource.ID, old *SecurityGroup, new *SecurityGroup, diff *resource.ObjectDiff) ([]string, error)
    Update(ctx context.Context,
        id resource.ID, old *SecurityGroup, new *SecurityGroup, diff *resource.ObjectDiff) error
    Delete(ctx context.Context, id resource.ID) error
}

// SecurityGroupProvider is a dynamic gRPC-based plugin for managing SecurityGroup resources.
type SecurityGroupProvider struct {
    ops SecurityGroupProviderOps
}

// NewSecurityGroupProvider allocates a resource provider that delegates to a ops instance.
func NewSecurityGroupProvider(ops SecurityGroupProviderOps) lumirpc.ResourceProviderServer {
    contract.Assert(ops != nil)
    return &SecurityGroupProvider{ops: ops}
}

func (p *SecurityGroupProvider) Check(
    ctx context.Context, req *lumirpc.CheckRequest) (*lumirpc.CheckResponse, error) {
    contract.Assert(req.GetType() == string(SecurityGroupToken))
    obj, _, err := p.Unmarshal(req.GetProperties())
    if err != nil {
        return plugin.NewCheckResponse(err), nil
    }
    if failures, err := p.ops.Check(ctx, obj); err != nil {
        return nil, err
    } else if len(failures) > 0 {
        return plugin.NewCheckResponse(resource.NewErrors(failures)), nil
    }
    return plugin.NewCheckResponse(nil), nil
}

func (p *SecurityGroupProvider) Name(
    ctx context.Context, req *lumirpc.NameRequest) (*lumirpc.NameResponse, error) {
    contract.Assert(req.GetType() == string(SecurityGroupToken))
    obj, _, err := p.Unmarshal(req.GetProperties())
    if err != nil {
        return nil, err
    }
    if obj.Name == nil || *obj.Name == "" {
        if req.Unknowns[SecurityGroup_Name] {
            return nil, errors.New("Name property cannot be computed from unknown outputs")
        }
        return nil, errors.New("Name property cannot be empty")
    }
    return &lumirpc.NameResponse{Name: *obj.Name}, nil
}

func (p *SecurityGroupProvider) Create(
    ctx context.Context, req *lumirpc.CreateRequest) (*lumirpc.CreateResponse, error) {
    contract.Assert(req.GetType() == string(SecurityGroupToken))
    obj, _, err := p.Unmarshal(req.GetProperties())
    if err != nil {
        return nil, err
    }
    id, err := p.ops.Create(ctx, obj)
    if err != nil {
        return nil, err
    }
    return &lumirpc.CreateResponse{Id: string(id)}, nil
}

func (p *SecurityGroupProvider) Get(
    ctx context.Context, req *lumirpc.GetRequest) (*lumirpc.GetResponse, error) {
    contract.Assert(req.GetType() == string(SecurityGroupToken))
    id := resource.ID(req.GetId())
    obj, err := p.ops.Get(ctx, id)
    if err != nil {
        return nil, err
    }
    return &lumirpc.GetResponse{
        Properties: plugin.MarshalProperties(
            nil, resource.NewPropertyMap(obj), plugin.MarshalOptions{}),
    }, nil
}

func (p *SecurityGroupProvider) InspectChange(
    ctx context.Context, req *lumirpc.InspectChangeRequest) (*lumirpc.InspectChangeResponse, error) {
    contract.Assert(req.GetType() == string(SecurityGroupToken))
    id := resource.ID(req.GetId())
    old, oldprops, err := p.Unmarshal(req.GetOlds())
    if err != nil {
        return nil, err
    }
    new, newprops, err := p.Unmarshal(req.GetNews())
    if err != nil {
        return nil, err
    }
    var replaces []string
    diff := oldprops.Diff(newprops)
    if diff != nil {
        if diff.Changed("name") {
            replaces = append(replaces, "name")
        }
        if diff.Changed("groupDescription") {
            replaces = append(replaces, "groupDescription")
        }
        if diff.Changed("groupName") {
            replaces = append(replaces, "groupName")
        }
        if diff.Changed("vpc") {
            replaces = append(replaces, "vpc")
        }
    }
    more, err := p.ops.InspectChange(ctx, id, old, new, diff)
    if err != nil {
        return nil, err
    }
    return &lumirpc.InspectChangeResponse{
        Replaces: append(replaces, more...),
    }, err
}

func (p *SecurityGroupProvider) Update(
    ctx context.Context, req *lumirpc.UpdateRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(SecurityGroupToken))
    id := resource.ID(req.GetId())
    old, oldprops, err := p.Unmarshal(req.GetOlds())
    if err != nil {
        return nil, err
    }
    new, newprops, err := p.Unmarshal(req.GetNews())
    if err != nil {
        return nil, err
    }
    diff := oldprops.Diff(newprops)
    if err := p.ops.Update(ctx, id, old, new, diff); err != nil {
        return nil, err
    }
    return &pbempty.Empty{}, nil
}

func (p *SecurityGroupProvider) Delete(
    ctx context.Context, req *lumirpc.DeleteRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(SecurityGroupToken))
    id := resource.ID(req.GetId())
    if err := p.ops.Delete(ctx, id); err != nil {
        return nil, err
    }
    return &pbempty.Empty{}, nil
}

func (p *SecurityGroupProvider) Unmarshal(
    v *pbstruct.Struct) (*SecurityGroup, resource.PropertyMap, error) {
    var obj SecurityGroup
    props := plugin.UnmarshalProperties(nil, v, plugin.MarshalOptions{RawResources: true})
    return &obj, props, mapper.MapIU(props.Mappable(), &obj)
}

/* Marshalable SecurityGroup structure(s) */

// SecurityGroup is a marshalable representation of its corresponding IDL type.
type SecurityGroup struct {
    Name *string `lumi:"name,optional"`
    GroupDescription string `lumi:"groupDescription"`
    GroupName *string `lumi:"groupName,optional"`
    VPC *resource.ID `lumi:"vpc,optional"`
    SecurityGroupEgress *[]SecurityGroupRule `lumi:"securityGroupEgress,optional"`
    SecurityGroupIngress *[]SecurityGroupRule `lumi:"securityGroupIngress,optional"`
    GroupID string `lumi:"groupID,optional"`
}

// SecurityGroup's properties have constants to make dealing with diffs and property bags easier.
const (
    SecurityGroup_Name = "name"
    SecurityGroup_GroupDescription = "groupDescription"
    SecurityGroup_GroupName = "groupName"
    SecurityGroup_VPC = "vpc"
    SecurityGroup_SecurityGroupEgress = "securityGroupEgress"
    SecurityGroup_SecurityGroupIngress = "securityGroupIngress"
    SecurityGroup_GroupID = "groupID"
)

/* Marshalable SecurityGroupRule structure(s) */

// SecurityGroupRule is a marshalable representation of its corresponding IDL type.
type SecurityGroupRule struct {
    IPProtocol string `lumi:"ipProtocol"`
    CIDRIP *string `lumi:"cidrIp,optional"`
    FromPort *float64 `lumi:"fromPort,optional"`
    ToPort *float64 `lumi:"toPort,optional"`
}

// SecurityGroupRule's properties have constants to make dealing with diffs and property bags easier.
const (
    SecurityGroupRule_IPProtocol = "ipProtocol"
    SecurityGroupRule_CIDRIP = "cidrIp"
    SecurityGroupRule_FromPort = "fromPort"
    SecurityGroupRule_ToPort = "toPort"
)


