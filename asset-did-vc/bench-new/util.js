'use strict';

function createRequest(operation, args) {
  const query = operation === 'query';
  return {
    contractId: 'did',
    contractVersion: '1.0',
    contractFunction: operation,
    contractArguments: Object.keys(args).map(k => args[k].toString()),
    readOnly: query
  };
}

module.exports.createRequest = createRequest;
