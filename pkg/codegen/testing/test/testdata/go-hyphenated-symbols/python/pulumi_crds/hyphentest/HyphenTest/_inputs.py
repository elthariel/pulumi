# coding=utf-8
# *** WARNING: this file was generated by crd2pulumi. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = [
    'HyphenTestSpecArgs',
]

@pulumi.input_type
class HyphenTestSpecArgs:
    def __init__(__self__, *,
                 has_a_hyphen: Optional[pulumi.Input[str]] = None):
        if has_a_hyphen is not None:
            pulumi.set(__self__, "has_a_hyphen", has_a_hyphen)

    @property
    @pulumi.getter(name="has-a-hyphen")
    def has_a_hyphen(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "has_a_hyphen")

    @has_a_hyphen.setter
    def has_a_hyphen(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "has_a_hyphen", value)


