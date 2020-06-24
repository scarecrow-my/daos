#!/usr/bin/python
"""
  (C) Copyright 2020 Intel Corporation.

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.

  GOVERNMENT LICENSE RIGHTS-OPEN SOURCE SOFTWARE
  The Government's rights to use, modify, reproduce, release, perform, display,
  or disclose this software are subject to the terms of the Apache License as
  provided in Contract No. B609815.
  Any reproduction of computer software, computer software documentation, or
  portions thereof marked with this legend must also reproduce the markings.
"""
from __future__ import print_function

from command_utils_base import \
    FormattedParameter, CommandWithParameters, YamlParameters
from command_utils import CommandWithSubCommand, YamlCommand


class DmgCommandBase(YamlCommand):
    """Defines a base object representing a dmg command."""

    METHOD_REGEX = {
        "run":
            r"(.*)",
        "network_scan":
            r"[-]+(?:\n|\n\r)([a-z0-9-]+)(?:\n|\n\r)[-]+|NUMA\s+"
            r"Socket\s+(\d+)|(ofi\+[a-z0-9;_]+)\s+([a-z0-9, ]+)",
        "pool_list":
            r"(?:([0-9a-fA-F-]+) +([0-9,]+))",
        "pool_create":
            r"(?:UUID:|Service replicas:)\s+([A-Za-z0-9-]+)",
        "pool_query":
            r"(?:Pool\s+([0-9a-fA-F-]+),\s+ntarget=(\d+),\s+disabled=(\d+),"
            r"\s+leader=(\d+),\s+version=(\d+)|Target\(VOS\)\s+count:"
            r"\s*(\d+)|(?:(?:SCM:|NVMe:)\s+Total\s+size:\s+([0-9.]+\s+[A-Z]+)"
            r"\s+Free:\s+([0-9.]+\s+[A-Z]+),\smin:([0-9.]+\s+[A-Z]+),"
            r"\s+max:([0-9.]+\s+[A-Z]+),\s+mean:([0-9.]+\s+[A-Z]+))"
            r"|Rebuild\s+\w+,\s+([0-9]+)\s+objs,\s+([0-9]+)\s+recs)",
        "storage_query_smd":
            r"(?:UUID|VOS\s+Target\s+IDs|SPDK Blobs):\s+([a-z0-9- ]+)",
        "storage_query_blobstore":
            r"(?:Device\s+UUID|Read\s+errors|Write\s+errors|Unmap\s+errors|"
            r"Checksum\s+errors|Error\s+log\s+entries|Media\s+errors|"
            r"Temperature|Available\s+Spare|Device\s+Reliability|"
            r"Read\s+Only|Volatile\s+Memory\s+Backup):\s+([A-Za-z0-9- ]+)",
        "storage_query_device_state":
            r"(?:Device\s+UUID|State):\s+([A-Za-z0-9-]+)",
        "storage_set_faulty":
            r"(?:Device\s+UUID|State):\s+([A-Za-z0-9-]+)",
        "system_query":
            r"(\d+|\[[0-9-,]+\])\s+([A-Za-z]+)",
        "system_start":
            r"(\d+|\[[0-9-,]+\])\s+([A-Za-z]+)\s+([A-Za-z]+)",
        "system_stop":
            r"(\d+|\[[0-9-,]+\])\s+([A-Za-z]+)\s+([A-Za-z]+)",
    }

    def __init__(self, path, yaml_cfg=None):
        """Create a dmg Command object.

        Args:
            path (str): path to the dmg command
            yaml_cfg (DmgYamlParameters, optional): dmg config file
                settings. Defaults to None, in which case settings
                must be supplied as command-line parameters.
        """
        super(DmgCommandBase, self).__init__(
            "/run/dmg/*", "dmg", path, yaml_cfg)

        # If specified use the configuration file from the YamlParameters object
        default_yaml_file = None
        if isinstance(self.yaml, YamlParameters):
            default_yaml_file = self.yaml.filename

        self._hostlist = FormattedParameter("-l {}")
        self.hostfile = FormattedParameter("-f {}")
        self.configpath = FormattedParameter("-o {}", default_yaml_file)
        self.insecure = FormattedParameter("-i", False)
        self.debug = FormattedParameter("-d", False)
        self.json = FormattedParameter("-j", False)

    @property
    def hostlist(self):
        """Get the hostlist that was set.

        Returns a string list.
        """
        if self.yaml:
            hosts = self.yaml.hostlist.value
        else:
            hosts = self._hostlist.value.split(",")
        return hosts

    @hostlist.setter
    def hostlist(self, hostlist):
        """Set the hostlist to be used for dmg invocation.

        Args:
            hostlist (string list): list of host addresses
        """
        if self.yaml:
            if not isinstance(hostlist, list):
                hostlist = hostlist.split(",")
            self.yaml.hostlist.update(hostlist, "dmg.yaml.hostlist")
        else:
            if isinstance(hostlist, list):
                hostlist = ",".join(hostlist)
            self._hostlist.update(hostlist, "dmg._hostlist")

    def get_sub_command_class(self):
        # pylint: disable=redefined-variable-type
        """Get the dmg sub command object based upon the sub-command."""
        if self.sub_command.value == "network":
            self.sub_command_class = self.NetworkSubCommand()
        elif self.sub_command.value == "pool":
            self.sub_command_class = self.PoolSubCommand()
        elif self.sub_command.value == "storage":
            self.sub_command_class = self.StorageSubCommand()
        elif self.sub_command.value == "system":
            self.sub_command_class = self.SystemSubCommand()
        else:
            self.sub_command_class = None

    class NetworkSubCommand(CommandWithSubCommand):
        """Defines an object for the dmg network sub command."""

        def __init__(self):
            """Create a dmg network subcommand object."""
            super(DmgCommandBase.NetworkSubCommand, self).__init__(
                "/run/dmg/network/*", "network")

        def get_sub_command_class(self):
            # pylint: disable=redefined-variable-type
            """Get the dmg network sub command object."""
            if self.sub_command.value == "scan":
                self.sub_command_class = self.ScanSubCommand()
            else:
                self.sub_command_class = None

        class ScanSubCommand(CommandWithParameters):
            """Defines an object for the dmg network scan command."""

            def __init__(self):
                """Create a dmg network scan command object."""
                super(
                    DmgCommandBase.NetworkSubCommand.ScanSubCommand,
                    self).__init__(
                        "/run/dmg/network/scan/*", "scan")
                self.provider = FormattedParameter("-p {}", None)
                self.all = FormattedParameter("-a", False)

    class PoolSubCommand(CommandWithSubCommand):
        """Defines an object for the dmg pool sub command."""

        def __init__(self):
            """Create a dmg pool subcommand object."""
            super(DmgCommandBase.PoolSubCommand, self).__init__(
                "/run/dmg/pool/*", "pool")

        def get_sub_command_class(self):
            # pylint: disable=redefined-variable-type
            """Get the dmg pool sub command object."""
            if self.sub_command.value == "create":
                self.sub_command_class = self.CreateSubCommand()
            elif self.sub_command.value == "delete-acl":
                self.sub_command_class = self.DeleteAclSubCommand()
            elif self.sub_command.value == "destroy":
                self.sub_command_class = self.DestroySubCommand()
            elif self.sub_command.value == "get-acl":
                self.sub_command_class = self.GetAclSubCommand()
            elif self.sub_command.value == "list":
                self.sub_command_class = self.ListSubCommand()
            elif self.sub_command.value == "overwrite-acl":
                self.sub_command_class = self.OverwriteAclSubCommand()
            elif self.sub_command.value == "query":
                self.sub_command_class = self.QuerySubCommand()
            elif self.sub_command.value == "set-prop":
                self.sub_command_class = self.SetPropSubCommand()
            elif self.sub_command.value == "update-acl":
                self.sub_command_class = self.UpdateAclSubCommand()
            elif self.sub_command.value == "exclude":
                self.sub_command_class = self.ExcludeSubCommand()
            elif self.sub_command.value == "reintegrate":
                self.sub_command_class = self.ReintegrateSubCommand()
            else:
                self.sub_command_class = None

        class CreateSubCommand(CommandWithParameters):
            """Defines an object for the dmg pool create command."""

            def __init__(self):
                """Create a dmg pool create command object."""
                super(
                    DmgCommandBase.PoolSubCommand.CreateSubCommand,
                    self).__init__(
                        "/run/dmg/pool/create/*", "create")
                self.group = FormattedParameter("--group={}", None)
                self.user = FormattedParameter("--user={}", None)
                self.acl_file = FormattedParameter("--acl-file={}", None)
                self.scm_size = FormattedParameter("--scm-size={}", None)
                self.nvme_size = FormattedParameter("--nvme-size={}", None)
                self.ranks = FormattedParameter("--ranks={}", None)
                self.nsvc = FormattedParameter("--nsvc={}", None)
                self.sys = FormattedParameter("--sys={}", None)

        class ExcludeSubCommand(CommandWithParameters):
            """Defines an object for the dmg pool exclude command."""

            def __init__(self):
                """Create a dmg pool exclude command object."""
                super(
                    DmgCommandBase.PoolSubCommand.ExcludeSubCommand,
                    self).__init__(
                        "/run/dmg/pool/exclude/*", "exclude")
                self.pool = FormattedParameter("--pool={}", None)
                self.rank = FormattedParameter("--rank={}", None)
                self.tgt_idx = FormattedParameter("--target-idx={}", None)

        class ReintegrateSubCommand(CommandWithParameters):
            """Defines an object for dmg pool reintegrate command."""

            def __init__(self):
                """Create a dmg pool reintegrate command object."""
                super(
                    DmgCommandBase.PoolSubCommand.ReintegrateSubCommand,
                    self).__init__(
                        "/run/dmg/pool/reintegrate/*", "reintegrate")
                self.pool = FormattedParameter("--pool={}", None)
                self.rank = FormattedParameter("--rank={}", None)
                self.tgt_idx = FormattedParameter("--target-idx={}", None)

        class DeleteAclSubCommand(CommandWithParameters):
            """Defines an object for the dmg pool delete-acl command."""

            def __init__(self):
                """Create a dmg pool delete-acl command object."""
                super(
                    DmgCommandBase.PoolSubCommand.DeleteAclSubCommand,
                    self).__init__(
                        "/run/dmg/pool/delete-acl/*", "delete-acl")
                self.pool = FormattedParameter("--pool={}", None)
                self.principal = FormattedParameter("-p {}", None)

        class DestroySubCommand(CommandWithParameters):
            """Defines an object for the dmg pool destroy command."""

            def __init__(self):
                """Create a dmg pool destroy command object."""
                super(
                    DmgCommandBase.PoolSubCommand.DestroySubCommand,
                    self).__init__(
                        "/run/dmg/pool/destroy/*", "destroy")
                self.pool = FormattedParameter("--pool={}", None)
                self.sys_name = FormattedParameter("--sys-name={}", None)
                self.force = FormattedParameter("--force", False)

        class GetAclSubCommand(CommandWithParameters):
            """Defines an object for the dmg pool get-acl command."""

            def __init__(self):
                """Create a dmg pool get-acl command object."""
                super(
                    DmgCommandBase.PoolSubCommand.GetAclSubCommand,
                    self).__init__(
                        "/run/dmg/pool/get-acl/*", "get-acl")
                self.pool = FormattedParameter("--pool={}", None)

        class ListSubCommand(CommandWithParameters):
            """Defines an object for the dmg pool list command."""

            def __init__(self):
                """Create a dmg pool list command object."""
                super(
                    DmgCommandBase.PoolSubCommand.ListSubCommand,
                    self).__init__(
                        "/run/dmg/pool/list/*", "list")

        class OverwriteAclSubCommand(CommandWithParameters):
            """Defines an object for the dmg pool overwrite-acl command."""

            def __init__(self):
                """Create a dmg pool overwrite-acl command object."""
                super(
                    DmgCommandBase.PoolSubCommand.OverwriteAclSubCommand,
                    self).__init__(
                        "/run/dmg/pool/overwrite-acl/*", "overwrite-acl")
                self.pool = FormattedParameter("--pool={}", None)
                self.acl_file = FormattedParameter("-a {}", None)

        class QuerySubCommand(CommandWithParameters):
            """Defines an object for the dmg pool query command."""

            def __init__(self):
                """Create a dmg pool query command object."""
                super(
                    DmgCommandBase.PoolSubCommand.QuerySubCommand,
                    self).__init__(
                        "/run/dmg/pool/query/*", "query")
                self.pool = FormattedParameter("--pool={}", None)

        class SetPropSubCommand(CommandWithParameters):
            """Defines an object for the dmg pool set-prop command."""

            def __init__(self):
                """Create a dmg pool set-prop command object."""
                super(
                    DmgCommandBase.PoolSubCommand.SetPropSubCommand,
                    self).__init__(
                        "/run/dmg/pool/set-prop/*", "set-prop")
                self.pool = FormattedParameter("--pool={}", None)
                self.name = FormattedParameter("--name={}", None)
                self.value = FormattedParameter("--value={}", None)

        class UpdateAclSubCommand(CommandWithParameters):
            """Defines an object for the dmg pool update-acl command."""

            def __init__(self):
                """Create a dmg pool update-acl command object."""
                super(
                    DmgCommandBase.PoolSubCommand.UpdateAclSubCommand,
                    self).__init__(
                        "/run/dmg/pool/update-acl/*", "update-acl")
                self.pool = FormattedParameter("--pool={}", None)
                self.acl_file = FormattedParameter("-a {}", None)
                self.entry = FormattedParameter("-e {}", None)

    class StorageSubCommand(CommandWithSubCommand):
        """Defines an object for the dmg storage sub command."""

        def __init__(self):
            """Create a dmg storage subcommand object."""
            super(DmgCommandBase.StorageSubCommand, self).__init__(
                "/run/dmg/storage/*", "storage")

        def get_sub_command_class(self):
            # pylint: disable=redefined-variable-type
            """Get the dmg storage sub command object."""
            if self.sub_command.value == "format":
                self.sub_command_class = self.FormatSubCommand()
            elif self.sub_command.value == "prepare":
                self.sub_command_class = self.PrepareSubCommand()
            elif self.sub_command.value == "query":
                self.sub_command_class = self.QuerySubCommand()
            elif self.sub_command.value == "scan":
                self.sub_command_class = self.ScanSubCommand()
            elif self.sub_command.value == "set":
                self.sub_command_class = self.SetSubCommand()
            else:
                self.sub_command_class = None

        class FormatSubCommand(CommandWithParameters):
            """Defines an object for the dmg storage format command."""

            def __init__(self):
                """Create a dmg storage format command object."""
                super(
                    DmgCommandBase.StorageSubCommand.FormatSubCommand,
                    self).__init__(
                        "/run/dmg/storage/format/*", "format")
                self.reformat = FormattedParameter("--reformat", False)

        class PrepareSubCommand(CommandWithParameters):
            """Defines an object for the dmg storage format command."""

            def __init__(self):
                """Create a dmg storage prepare command object."""
                super(
                    DmgCommandBase.StorageSubCommand.PrepareSubCommand,
                    self).__init__(
                        "/run/dmg/storage/prepare/*", "prepare")
                self.pci_whitelist = FormattedParameter("-w {}", None)
                self.hugepages = FormattedParameter("-p {}", None)
                self.target_user = FormattedParameter("-u {}", None)
                self.nvme_only = FormattedParameter("-n", False)
                self.scm_only = FormattedParameter("-s", False)
                self.reset = FormattedParameter("--reset", False)
                self.force = FormattedParameter("-f", False)

        class QuerySubCommand(CommandWithSubCommand):
            """Defines an object for the dmg query format command."""

            def __init__(self):
                """Create a dmg storage query command object."""
                super(
                    DmgCommandBase.StorageSubCommand.QuerySubCommand,
                    self).__init__(
                        "/run/dmg/storage/query/*", "query")

            def get_sub_command_class(self):
                # pylint: disable=redefined-variable-type
                """Get the dmg pool sub command object."""
                if self.sub_command.value == "blobstore-health":
                    self.sub_command_class = self.BlobstoreHealthSubCommand()
                elif self.sub_command.value == "device-state":
                    self.sub_command_class = self.DeviceStateSubCommand()
                elif self.sub_command.value == "nvme-health":
                    self.sub_command_class = self.NvmeHealthSubCommand()
                elif self.sub_command.value == "smd":
                    self.sub_command_class = self.SmdSubCommand()
                else:
                    self.sub_command_class = None

            class BlobstoreHealthSubCommand(CommandWithParameters):
                """Defines a dmg storage query blobstore-health object."""

                def __init__(self):
                    """Create a dmg storage query blobstore-health object."""
                    super(
                        DmgCommandBase.StorageSubCommand.QuerySubCommand.
                        BlobstoreHealthSubCommand,
                        self).__init__(
                            "/run/dmg/storage/query/blobstore-health/*",
                            "blobstore-health")
                    self.devuuid = FormattedParameter("-u {}", None)
                    self.tgtid = FormattedParameter("-t {}", None)

            class DeviceStateSubCommand(CommandWithParameters):
                """Defines a dmg storage query device-state object."""

                def __init__(self):
                    """Create a dmg storage query device-state object."""
                    super(
                        DmgCommandBase.StorageSubCommand.QuerySubCommand.
                        DeviceStateSubCommand,
                        self).__init__(
                            "/run/dmg/storage/query/device-state/*",
                            "device-state")
                    self.devuuid = FormattedParameter("-u {}", None)

            class NvmeHealthSubCommand(CommandWithParameters):
                """Defines a dmg storage query nvme-health object."""

                def __init__(self):
                    """Create a dmg storage query nvme-health object."""
                    super(
                        DmgCommandBase.StorageSubCommand.QuerySubCommand.
                        NvmeHealthSubCommand,
                        self).__init__(
                            "/run/dmg/storage/query/nvme-health/*",
                            "nvme-health")

            class SmdSubCommand(CommandWithParameters):
                """Defines a dmg storage query smd object."""

                def __init__(self):
                    """Create a dmg storage query smd object."""
                    super(
                        DmgCommandBase.StorageSubCommand.QuerySubCommand.
                        SmdSubCommand,
                        self).__init__(
                            "/run/dmg/storage/query/smd/*",
                            "smd")
                    self.devices = FormattedParameter("-d", False)
                    self.pools = FormattedParameter("-p", False)

        class ScanSubCommand(CommandWithParameters):
            """Defines an object for the dmg storage scan command."""

            def __init__(self):
                """Create a dmg storage scan command object."""
                super(
                    DmgCommandBase.StorageSubCommand.ScanSubCommand,
                    self).__init__(
                        "/run/dmg/storage/scan/*", "scan")
                self.summary = FormattedParameter("-m", False)
                self.verbose = FormattedParameter("--verbose", False)

        class SetSubCommand(CommandWithSubCommand):
            """Defines an object for the dmg storage set command."""

            def __init__(self):
                """Create a dmg storage set command object."""
                super(
                    DmgCommandBase.StorageSubCommand.SetSubCommand,
                    self).__init__(
                        "/run/dmg/storage/set/*", "set")

            def get_sub_command_class(self):
                # pylint: disable=redefined-variable-type
                """Get the dmg set sub command object."""
                if self.sub_command.value == "nvme-faulty":
                    self.sub_command_class = self.NvmeFaultySubCommand()
                else:
                    self.sub_command_class = None

            class NvmeFaultySubCommand(CommandWithParameters):
                """Defines a dmg storage set nvme-faulty object."""

                def __init__(self):
                    """Create a dmg storage set nvme-faulty object."""
                    super(
                        DmgCommandBase.StorageSubCommand.SetSubCommand.
                        NvmeFaultySubCommand,
                        self).__init__(
                            "/run/dmg/storage/query/device-state/*",
                            "nvme-faulty")
                    self.devuuid = FormattedParameter("-u {}", None)

    class SystemSubCommand(CommandWithSubCommand):
        """Defines an object for the dmg system sub command."""

        def __init__(self):
            """Create a dmg system subcommand object."""
            super(DmgCommandBase.SystemSubCommand, self).__init__(
                "/run/dmg/system/*", "system")

        def get_sub_command_class(self):
            # pylint: disable=redefined-variable-type
            """Get the dmg system sub command object."""
            if self.sub_command.value == "leader-query":
                self.sub_command_class = self.LeaderQuerySubCommand()
            elif self.sub_command.value == "list-pools":
                self.sub_command_class = self.ListPoolsSubCommand()
            elif self.sub_command.value == "query":
                self.sub_command_class = self.QuerySubCommand()
            elif self.sub_command.value == "start":
                self.sub_command_class = self.StartSubCommand()
            elif self.sub_command.value == "stop":
                self.sub_command_class = self.StopSubCommand()
            else:
                self.sub_command_class = None

        class LeaderQuerySubCommand(CommandWithParameters):
            """Defines an object for the dmg system leader-query command."""

            def __init__(self):
                """Create a dmg system leader-query command object."""
                super(
                    DmgCommandBase.SystemSubCommand.LeaderQuerySubCommand,
                    self).__init__(
                        "/run/dmg/system/leader-query/*", "leader-query")

        class ListPoolsSubCommand(CommandWithParameters):
            """Defines an object for the dmg system list-pools command."""

            def __init__(self):
                """Create a dmg system list-pools command object."""
                super(
                    DmgCommandBase.SystemSubCommand.ListPoolsSubCommand,
                    self).__init__(
                        "/run/dmg/system/list-pools/*", "list-pools")

        class QuerySubCommand(CommandWithParameters):
            """Defines an object for the dmg system query command."""

            def __init__(self):
                """Create a dmg system query command object."""
                super(
                    DmgCommandBase.SystemSubCommand.QuerySubCommand,
                    self).__init__(
                        "/run/dmg/system/query/*", "query")
                self.rank = FormattedParameter("--rank={}")
                self.verbose = FormattedParameter("--verbose", False)

        class StartSubCommand(CommandWithParameters):
            """Defines an object for the dmg system start command."""

            def __init__(self):
                """Create a dmg system start command object."""
                super(
                    DmgCommandBase.SystemSubCommand.StartSubCommand,
                    self).__init__(
                        "/run/dmg/system/start/*", "start")

        class StopSubCommand(CommandWithParameters):
            """Defines an object for the dmg system stop command."""

            def __init__(self):
                """Create a dmg system stop command object."""
                super(
                    DmgCommandBase.SystemSubCommand.StopSubCommand,
                    self).__init__(
                        "/run/dmg/system/stop/*", "stop")
                self.force = FormattedParameter("--force", False)