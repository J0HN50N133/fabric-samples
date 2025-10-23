/*
* Licensed under the Apache License, Version 2.0 (the \"License\");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
* http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an \"AS IS\" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*/

'use strict';

const { createConnectorRequest } = require('./util.js');
const { WorkloadModuleBase } = require('@hyperledger/caliper-core');

/**
 * Workload module for querying the accumulator.
 */
class GetAccumulator extends WorkloadModuleBase {

  /**
   * Initializes the parameters of the workload.
   */
  constructor() {
    super();
  }

  /**
   * Assemble TXs for querying the accumulator.
   */
  async submitTransaction() {
    const myArgs = {};
    await this.sutAdapter.sendRequests(createConnectorRequest('GetAccumulator', myArgs));
  }
}

/**
 * Create a new instance of the workload module.
 * @return {WorkloadModuleInterface}
 */
function createWorkloadModule() {
  return new GetAccumulator();
}

module.exports.createWorkloadModule = createWorkloadModule;

