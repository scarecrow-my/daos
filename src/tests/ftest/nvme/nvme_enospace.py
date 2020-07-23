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
from nvme_utils import ServerFillUp
from avocado.core.exceptions import TestFail
from general_utils import get_log_file, run_task

class NvmeEnospace(ServerFillUp):
    # pylint: disable=too-many-ancestors
    """
    Test Class Description: To validate DER_NOSPACE for SCM and NVMe
    :avocado: recursive
    """
    def der_enspace_log_count(self):
        """
        Function to count the DER_NOSPACE and other ERR in client log.

        arg:
            None

        returns:
            der_nospace_count(int): DER_NOSPACE count from client log.
            other_errors_count(int): Other Error count from client log.

        """
        #Get the Client side Error from client_log file.
        cmd = 'cat {} | grep ERR'.format(get_log_file(self.client_log))
        task = run_task(self.hostlist_clients, cmd)
        for _rc_code, _node in task.iter_retcodes():
            if _rc_code == 1:
                self.fail("Failed to run cmd {} on {}".format(cmd, _node))
        for buf, _nodes in task.iter_buffers():
            output = str(buf).split('\n')

        der_nospace_count = 0
        other_errors_count = 0
        for line in output:
            if 'DER_NOSPACE' in line:
                der_nospace_count += 1
            else:
                other_errors_count += 1
        print('Total Error={} DER_NOSPACE Error={} Other Error={}'
              .format(len(output), der_nospace_count, other_errors_count))

        return der_nospace_count, other_errors_count

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

        self.create_pool_max_size()
        print(self.pool.pool_percentage_used())

        # Disable the aggregation
        self.pool.set_property("reclaim", "disabled")

        #Fill 75% of SCM pool
        self.start_ior_load(storage='SCM', precent=75)

        print(self.container_info)
        print(self.pool.pool_percentage_used())

        try:
            #Fill 10% more to SCM ,which suppose to Fail because no SCM space
            self.start_ior_load(storage='SCM', precent=10)
            ##self.fail('This test suppose to Fail but it got Passed')
        except TestFail as _error:
            self.log.info('Test should get failed as expected')

        der_count, other_error = self.der_enspace_log_count()
        if other_error > 0:
            self.fail('Found other count {} in client log {}'
                      .format(other_error, self.client_log))
        if der_count != 1:
            self.fail('Expected DER_NOSPACE should be 1 and Found {}'
                      .format(der_count))

        #Destroy the container and loop through again
        print(self.container_info)
        self.pool.destroy()
