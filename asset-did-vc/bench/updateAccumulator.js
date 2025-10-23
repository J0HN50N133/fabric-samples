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

const OperationBase = require('../../caliper-benchmarks/benchmarks/scenario/simple/utils/operation-base');
const crypto = require('crypto');

/**
 * Workload module for updating the accumulator.
 */
class UpdateAccumulator extends OperationBase {

    /**
     * Initializes the parameters of the workload.
     */
    constructor() {
        super();
    }

    /**
     * Assemble TXs for updating the accumulator.
     */
    async submitTransaction() {
        // Generate a random 32-byte hex string for the new digest
        const newDigest = crypto.randomBytes(32).toString('hex');
        const myArgs = {
            newDigestHex: newDigest
        };
        await this.sutAdapter.sendRequests(this.createConnectorRequest('UpdateAccumulator', myArgs));
    }
}

/**
 * Create a new instance of the workload module.
 * @return {WorkloadModuleInterface}
 */
function createWorkloadModule() {
    return new UpdateAccumulator();
}

module.exports.createWorkloadModule = createWorkloadModule;

