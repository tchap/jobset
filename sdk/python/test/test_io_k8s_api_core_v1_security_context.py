# coding: utf-8

"""
    JobSet SDK

    Python SDK for the JobSet API

    The version of the OpenAPI document: v0.1.4
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


import unittest

from jobset.models.io_k8s_api_core_v1_security_context import IoK8sApiCoreV1SecurityContext

class TestIoK8sApiCoreV1SecurityContext(unittest.TestCase):
    """IoK8sApiCoreV1SecurityContext unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> IoK8sApiCoreV1SecurityContext:
        """Test IoK8sApiCoreV1SecurityContext
            include_optional is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `IoK8sApiCoreV1SecurityContext`
        """
        model = IoK8sApiCoreV1SecurityContext()
        if include_optional:
            return IoK8sApiCoreV1SecurityContext(
                allow_privilege_escalation = True,
                app_armor_profile = jobset.models.io/k8s/api/core/v1/app_armor_profile.io.k8s.api.core.v1.AppArmorProfile(
                    localhost_profile = '', 
                    type = '', ),
                capabilities = jobset.models.io/k8s/api/core/v1/capabilities.io.k8s.api.core.v1.Capabilities(
                    add = [
                        ''
                        ], 
                    drop = [
                        ''
                        ], ),
                privileged = True,
                proc_mount = '',
                read_only_root_filesystem = True,
                run_as_group = 56,
                run_as_non_root = True,
                run_as_user = 56,
                se_linux_options = jobset.models.io/k8s/api/core/v1/se_linux_options.io.k8s.api.core.v1.SELinuxOptions(
                    level = '', 
                    role = '', 
                    type = '', 
                    user = '', ),
                seccomp_profile = jobset.models.io/k8s/api/core/v1/seccomp_profile.io.k8s.api.core.v1.SeccompProfile(
                    localhost_profile = '', 
                    type = '', ),
                windows_options = jobset.models.io/k8s/api/core/v1/windows_security_context_options.io.k8s.api.core.v1.WindowsSecurityContextOptions(
                    gmsa_credential_spec = '', 
                    gmsa_credential_spec_name = '', 
                    host_process = True, 
                    run_as_user_name = '', )
            )
        else:
            return IoK8sApiCoreV1SecurityContext(
        )
        """

    def testIoK8sApiCoreV1SecurityContext(self):
        """Test IoK8sApiCoreV1SecurityContext"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
