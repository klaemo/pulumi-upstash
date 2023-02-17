# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['QStashEndpointArgs', 'QStashEndpoint']

@pulumi.input_type
class QStashEndpointArgs:
    def __init__(__self__, *,
                 topic_id: pulumi.Input[str],
                 url: pulumi.Input[str]):
        """
        The set of arguments for constructing a QStashEndpoint resource.
        :param pulumi.Input[str] topic_id: Topic Id that the endpoint is added to
        :param pulumi.Input[str] url: URL of the endpoint
        """
        pulumi.set(__self__, "topic_id", topic_id)
        pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter(name="topicId")
    def topic_id(self) -> pulumi.Input[str]:
        """
        Topic Id that the endpoint is added to
        """
        return pulumi.get(self, "topic_id")

    @topic_id.setter
    def topic_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "topic_id", value)

    @property
    @pulumi.getter
    def url(self) -> pulumi.Input[str]:
        """
        URL of the endpoint
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: pulumi.Input[str]):
        pulumi.set(self, "url", value)


@pulumi.input_type
class _QStashEndpointState:
    def __init__(__self__, *,
                 endpoint_id: Optional[pulumi.Input[str]] = None,
                 topic_id: Optional[pulumi.Input[str]] = None,
                 topic_name: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering QStashEndpoint resources.
        :param pulumi.Input[str] endpoint_id: Unique Qstash Endpoint ID
        :param pulumi.Input[str] topic_id: Topic Id that the endpoint is added to
        :param pulumi.Input[str] topic_name: Unique Qstash Topic Name for Endpoint
        :param pulumi.Input[str] url: URL of the endpoint
        """
        if endpoint_id is not None:
            pulumi.set(__self__, "endpoint_id", endpoint_id)
        if topic_id is not None:
            pulumi.set(__self__, "topic_id", topic_id)
        if topic_name is not None:
            pulumi.set(__self__, "topic_name", topic_name)
        if url is not None:
            pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter(name="endpointId")
    def endpoint_id(self) -> Optional[pulumi.Input[str]]:
        """
        Unique Qstash Endpoint ID
        """
        return pulumi.get(self, "endpoint_id")

    @endpoint_id.setter
    def endpoint_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "endpoint_id", value)

    @property
    @pulumi.getter(name="topicId")
    def topic_id(self) -> Optional[pulumi.Input[str]]:
        """
        Topic Id that the endpoint is added to
        """
        return pulumi.get(self, "topic_id")

    @topic_id.setter
    def topic_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "topic_id", value)

    @property
    @pulumi.getter(name="topicName")
    def topic_name(self) -> Optional[pulumi.Input[str]]:
        """
        Unique Qstash Topic Name for Endpoint
        """
        return pulumi.get(self, "topic_name")

    @topic_name.setter
    def topic_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "topic_name", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        URL of the endpoint
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)


class QStashEndpoint(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 topic_id: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import upstash_pulumi as upstash

        example_qstash_endpoint = upstash.QStashEndpoint("exampleQstashEndpoint",
            url="https://***.***",
            topic_id=resource["upstash_qstash_topic"]["exampleQstashTopic"]["topic_id"])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] topic_id: Topic Id that the endpoint is added to
        :param pulumi.Input[str] url: URL of the endpoint
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: QStashEndpointArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import upstash_pulumi as upstash

        example_qstash_endpoint = upstash.QStashEndpoint("exampleQstashEndpoint",
            url="https://***.***",
            topic_id=resource["upstash_qstash_topic"]["exampleQstashTopic"]["topic_id"])
        ```

        :param str resource_name: The name of the resource.
        :param QStashEndpointArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(QStashEndpointArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 topic_id: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = QStashEndpointArgs.__new__(QStashEndpointArgs)

            if topic_id is None and not opts.urn:
                raise TypeError("Missing required property 'topic_id'")
            __props__.__dict__["topic_id"] = topic_id
            if url is None and not opts.urn:
                raise TypeError("Missing required property 'url'")
            __props__.__dict__["url"] = url
            __props__.__dict__["endpoint_id"] = None
            __props__.__dict__["topic_name"] = None
        super(QStashEndpoint, __self__).__init__(
            'upstash:index/qStashEndpoint:QStashEndpoint',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            endpoint_id: Optional[pulumi.Input[str]] = None,
            topic_id: Optional[pulumi.Input[str]] = None,
            topic_name: Optional[pulumi.Input[str]] = None,
            url: Optional[pulumi.Input[str]] = None) -> 'QStashEndpoint':
        """
        Get an existing QStashEndpoint resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] endpoint_id: Unique Qstash Endpoint ID
        :param pulumi.Input[str] topic_id: Topic Id that the endpoint is added to
        :param pulumi.Input[str] topic_name: Unique Qstash Topic Name for Endpoint
        :param pulumi.Input[str] url: URL of the endpoint
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _QStashEndpointState.__new__(_QStashEndpointState)

        __props__.__dict__["endpoint_id"] = endpoint_id
        __props__.__dict__["topic_id"] = topic_id
        __props__.__dict__["topic_name"] = topic_name
        __props__.__dict__["url"] = url
        return QStashEndpoint(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="endpointId")
    def endpoint_id(self) -> pulumi.Output[str]:
        """
        Unique Qstash Endpoint ID
        """
        return pulumi.get(self, "endpoint_id")

    @property
    @pulumi.getter(name="topicId")
    def topic_id(self) -> pulumi.Output[str]:
        """
        Topic Id that the endpoint is added to
        """
        return pulumi.get(self, "topic_id")

    @property
    @pulumi.getter(name="topicName")
    def topic_name(self) -> pulumi.Output[str]:
        """
        Unique Qstash Topic Name for Endpoint
        """
        return pulumi.get(self, "topic_name")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        """
        URL of the endpoint
        """
        return pulumi.get(self, "url")
