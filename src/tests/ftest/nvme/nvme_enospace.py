#!/usr/bin/python
'''
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
'''
import os
from nvme_utils import ServerFillUp, get_device_ids
from test_utils_pool import TestPool
from dmg_utils import DmgCommand
from command_utils_base import CommandFailure

class NvmeHealth(ServerFillUp):
    # pylint: disable=too-many-ancestors
    """
    Test Class Description: To validate NVMe health test cases
    :avocado: recursive
    """
    def test_monitor_for_large_pools(self):
        """Jira ID: DAOS-4722.

        Test Description: Test Health monitor for large number of pools.
        Use Case: This tests will create the 40 number of pools and verify the
                  dmg list-pools, device-health and nvme-health works for all
                  pools.

        :avocado: tags=all,hw,medium,nvme,ib2,full_regression
        :avocado: tags=nvme_enospc
        """
        # pylint: disable=attribute-defined-outside-init
        # pylint: disable=too-many-branches

        #pool = TestPool(self.context, dmg_command=self.get_dmg_command())
        #pool.get_params(self)
        #pool.create()
        self.create_pool_max_size()
        print (self.pool.pool_percentage_used())

        # Disable the aggregation
        self.pool.set_property("reclaim", "disabled")

        #Fill % of SCM pool
        self.start_ior_load(storage='SCM', precent=1)

        print (self.pool.pool_percentage_used())
        self.pool.destroy()