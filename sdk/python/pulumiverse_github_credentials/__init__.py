# coding=utf-8
# *** WARNING: this file was generated by Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .provider import *
from .static_page import *
_utilities.register(
    resource_modules="""
[
 {
  "pkg": "github-credentials",
  "mod": "index",
  "fqn": "pulumiverse_github_credentials",
  "classes": {
   "github-credentials:index:StaticPage": "StaticPage"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "github-credentials",
  "token": "pulumi:providers:github-credentials",
  "fqn": "pulumiverse_github_credentials",
  "class": "Provider"
 }
]
"""
)