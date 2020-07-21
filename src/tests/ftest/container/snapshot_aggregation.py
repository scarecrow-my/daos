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
from ior_test_base import IorTestBase
from daos_utils import DaosCommand


class SnapshotAggregation(IorTestBase):
    # pylint: disable=too-many-ancestors
    """Defines snapshot aggregation test cases.

    Test Class Description:
          Verify snapshot aggregation.

    :avocado: recursive
    """

    def __init__(self, *args, **kwargs):
        """Initialize a SnapshotAggregation object."""
        super(SnapshotAggregation, self).__init__(*args, **kwargs)
        self.dmg = None
        self.free_space = {"SCM": [], "NVMe": []}

    def update_free_space(self):
        """Append the free space list with the current pool capacities."""
        self.pool.get_info()
        self.free_space["SCM"].append(self.pool.info.scm.free)
        self.free_space["NVMe"].append(self.pool.info.nvme.free)

        # pool_query = self.dmg.get_output("pool_query", pool=self.pool.uuid)
        # # When called through get_output() the 'dmg pool query' output is
        # # returned as a tuple of tuples after running through the regex.
        # #
        # # For example, running the regex against this output:
        # #   Pool <A>, ntarget=<B>, disabled=<C>, leader=<D>, version=<E>
        # #   Pool space info:
        # #   - Target(VOS) count:<F>
        # #   - SCM:
        # #     Total size: <G>
        # #     Free: <H>, min:<I>, max:<J>, mean:<K>
        # #   - NVMe:
        # #     Total size: <L>
        # #     Free: <M>, min:<N>, max:<O>, mean:<P>
        # #   Rebuild idle, <Q> objs, <R> recs
        # #
        # # Yields this tuple of tuples:
        # #   [0]: (<A>, <B>, <C>, <D>, <E>, '', '', '', '', '', '', '', '')
        # #   [1]: ('', '', '', '', '', <F>, '', '', '', '', '', '', '')
        # #   [2]: ('', '', '', '', '', '', <G>, <H>, <I>, <J>, <K>, '', '')
        # #   [3]: ('', '', '', '', '', '', <L>, <M>, <N>, <O>, <P>, '', '')
        # #   [4]: ('', '', '', '', '', '', '', '', '', '', '', <Q>, <R>)
        # #
        # # For this test we want the SCM and NVMe Free space: <H> and <M>.

        # for index, value in enumerate(pool_query):
        #     self.log.info("  [%s]: %s", index, value)
        #     for index2, value2 in enumerate(value):
        #         self.log.info("    [%s]: %s", index2, value2)

        # try:
        #     free_scm = pool_query[2][7].split(" ")
        #     free_nvme = pool_query[3][7].split(" ")
        # except IndexError as error:
        #     self.fail(
        #         "Error obtaining SCM and NVMe free space from dmg pool "
        #         "query: {}".format(error))

        # # Since sizes are not guaranteed to be in the same units, we'll use the
        # # Bytes class to be able to compare items.  As such, units like 'MB'
        # # need to be provided as 'M'.
        # try:
        #     self.free_space["SCM"].append(Bytes(free_scm[0], free_scm[1][0]))
        #     self.free_space["NVMe"].append(Bytes(free_nvme[0], free_nvme[1][0]))
        # except IndexError as error:
        #     self.fail(
        #         "Error assigning SCM and NVMe free space from '{}' and '{}', "
        #         "respectively: {}".format(free_scm, free_nvme, error))

    def test_snapshot_aggregation(self):
        """JIRA ID: DAOS-3751.

        Test Description:
            Verify snapshot aggregation with 2 servers and 6 clients (CI limit).
            Write the same data to the pool twice.  Create a snapshot between
            the writes and confirm that deleting the snapshot reduces the pool
            capacity by half.

        :avocado: tags=all,pr,hw,large,container,snapshot,snapshot_aggregation
        """
        self.dmg = self.get_dmg_command()
        daos = DaosCommand(self.bin)

        # Create a pool and a container that spans the 2 servers.
        self.update_ior_cmd_with_pool()
        self.update_free_space()
        self.log.info(
            "Pool free space before writes:\n  SCM:  %s\n  NVMe: %s",
            self.free_space["SCM"][-1], self.free_space["NVMe"][-1])

        # Run an IOR job that writes >4k sequential blocks for a few minutes
        self.processes = len(self.hostlist_clients)
        manager = self.get_ior_job_manager_command()
        self.run_ior(manager, self.processes)

        # Get the capacity of the pool after running IOR.
        self.update_free_space()
        self.log.info(
            "Pool free space after first write:\n  SCM:  %s\n  NVMe: %s",
            self.free_space["SCM"][-1], self.free_space["NVMe"][-1])

        # Create a snapshot of the container once the IOR job completes.
        snapshot = daos.get_output(
            "container_create_snap", pool=self.pool.uuid,
            svc=self.pool.svc_ranks, cont=self.container.uuid)

        # Run the same IOR job to cause an overwrite.
        self.ior_cmd.signature.value += 333
        self.run_ior(self.get_ior_job_manager_command(), self.processes)

        # Watch the DAOS logs looking for aggregation to start.

        # Wait to see aggregation completes by checking for a DAOS log message.

        # Verify that the utilized capacity of the pool has doubled.
        self.update_free_space()
        self.log.info(
            "Pool free space after second write:\n  SCM:  %s\n  NVMe: %s",
            self.free_space["SCM"][-1], self.free_space["NVMe"][-1])

        # Delete the snapshot.
        daos.container_destroy_snap(
            pool=self.pool.uuid, svc=self.pool.svc_ranks,
            cont=self.container.uuid, epc=snapshot)

        # Watch for aggregation to start & finish.

        # Verify that the utilized capacity of the pool has been halved.
        self.update_free_space()
        self.log.info(
            "Pool free space after deleting snapshot:\n  SCM:  %s\n  NVMe: %s",
            self.free_space["SCM"][-1], self.free_space["NVMe"][-1])