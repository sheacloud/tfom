/* eslint-disable */
import * as types from './graphql';
import { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';

/**
 * Map of all GraphQL operations in the project.
 *
 * This map has several performance disadvantages:
 * 1. It is not tree-shakeable, so it will include all operations in the project.
 * 2. It is not minifiable, so the string of a GraphQL query will be multiple times inside the bundle.
 * 3. It does not support dead code elimination, so it will add unused operations.
 *
 * Therefore it is highly recommended to use the babel-plugin for production.
 */
const documents = {
    "\n  query applyExecutionRequest($applyExecutionRequestId: ID!) {\n    applyExecutionRequest(applyExecutionRequestId: $applyExecutionRequestId) {\n      applyExecutionRequestId\n      status\n      requestTime\n      terraformConfigurationBase64\n      initOutput {\n        Stdout\n        Stderr\n      }\n      applyOutput {\n        Stdout\n        Stderr\n      }\n    }\n  }\n": types.ApplyExecutionRequestDocument,
    "\n  query moduleAccountAssociation(\n    $modulePropagationId: ID!\n    $orgAccountId: ID!\n  ) {\n    moduleAccountAssociation(\n      modulePropagationId: $modulePropagationId\n      orgAccountId: $orgAccountId\n    ) {\n      modulePropagationId\n      modulePropagation {\n        name\n        moduleGroup {\n          moduleGroupId\n          name\n        }\n        moduleVersion {\n          moduleVersionId\n          name\n        }\n      }\n      orgAccount {\n        orgAccountId\n        name\n      }\n      status\n      planExecutionRequests {\n        items {\n          planExecutionRequestId\n          status\n          requestTime\n          modulePropagationId\n          modulePropagationRequestId\n        }\n      }\n      applyExecutionRequests {\n        items {\n          applyExecutionRequestId\n          status\n          requestTime\n          modulePropagationId\n          modulePropagationRequestId\n        }\n      }\n    }\n  }\n": types.ModuleAccountAssociationDocument,
    "\n  query moduleGroups {\n    moduleGroups(limit: 100) {\n      items {\n        moduleGroupId\n        name\n      }\n    }\n  }\n": types.ModuleGroupsDocument,
    "\n  query moduleGroup($moduleGroupId: ID!) {\n    moduleGroup(moduleGroupId: $moduleGroupId) {\n      moduleGroupId\n      name\n      versions {\n        items {\n          moduleVersionId\n          name\n          remoteSource\n          terraformVersion\n        }\n      }\n      modulePropagations {\n        items {\n          name\n          modulePropagationId\n          moduleVersion {\n            moduleVersionId\n            name\n          }\n          orgUnit {\n            orgUnitId\n            name\n          }\n          orgDimension {\n            orgDimensionId\n            name\n          }\n        }\n      }\n    }\n  }\n": types.ModuleGroupDocument,
    "\n  query modulePropagationExecutionRequest(\n    $modulePropagationId: ID!\n    $modulePropagationExecutionRequestId: ID!\n  ) {\n    modulePropagationExecutionRequest(\n      modulePropagationId: $modulePropagationId\n      modulePropagationExecutionRequestId: $modulePropagationExecutionRequestId\n    ) {\n      modulePropagationId\n      modulePropagationExecutionRequestId\n      requestTime\n      status\n      terraformExecutionWorkflowRequests {\n        items {\n          terraformExecutionWorkflowRequestId\n          status\n          requestTime\n          destroy\n          orgAccount {\n            orgAccountId\n            name\n          }\n          planExecutionRequest {\n            planExecutionRequestId\n            status\n            requestTime\n          }\n          applyExecutionRequest {\n            applyExecutionRequestId\n            status\n            requestTime\n          }\n        }\n      }\n      planExecutionRequests {\n        items {\n          planExecutionRequestId\n          status\n          requestTime\n          orgAccount {\n            orgAccountId\n            name\n          }\n        }\n      }\n      applyExecutionRequests {\n        items {\n          applyExecutionRequestId\n          status\n          requestTime\n          orgAccount {\n            orgAccountId\n            name\n          }\n        }\n      }\n    }\n  }\n": types.ModulePropagationExecutionRequestDocument,
    "\n  query modulePropagation($modulePropagationId: ID!) {\n    modulePropagation(modulePropagationId: $modulePropagationId) {\n      modulePropagationId\n      moduleGroup {\n        moduleGroupId\n        name\n      }\n      moduleVersion {\n        moduleVersionId\n        name\n      }\n      orgUnitId\n      orgUnit {\n        orgUnitId\n        orgDimensionId\n        name\n        downstreamOrgUnits {\n          items {\n            orgUnitId\n            orgDimensionId\n            name\n          }\n        }\n      }\n      modulePropagationId\n      name\n      description\n      executionRequests(limit: 5) {\n        items {\n          modulePropagationId\n          modulePropagationExecutionRequestId\n          requestTime\n          status\n        }\n      }\n      moduleAccountAssociations {\n        items {\n          modulePropagationId\n          orgAccount {\n            orgAccountId\n            name\n          }\n          status\n        }\n      }\n    }\n  }\n": types.ModulePropagationDocument,
    "\n  mutation createModulePropagationExecutionRequest($modulePropagationId: ID!) {\n    createModulePropagationExecutionRequest(\n      modulePropagationExecutionRequest: {\n        modulePropagationId: $modulePropagationId\n      }\n    ) {\n      modulePropagationExecutionRequestId\n      status\n    }\n  }\n": types.CreateModulePropagationExecutionRequestDocument,
    "\n  query orgDimensions {\n    organizationalDimensions(limit: 10000) {\n      items {\n        orgDimensionId\n        name\n        orgUnits(limit: 10000) {\n          items {\n            orgUnitId\n            name\n          }\n        }\n      }\n    }\n  }\n": types.OrgDimensionsDocument,
    "\n  mutation updateModulePropagation(\n    $modulePropagationId: ID!\n    $update: ModulePropagationUpdate!\n  ) {\n    updateModulePropagation(\n      modulePropagationId: $modulePropagationId\n      update: $update\n    ) {\n      modulePropagationId\n    }\n  }\n": types.UpdateModulePropagationDocument,
    "\n  query moduleVersion($moduleGroupId: ID!, $moduleVersionId: ID!) {\n    moduleVersion(\n      moduleGroupId: $moduleGroupId\n      moduleVersionId: $moduleVersionId\n    ) {\n      moduleVersionId\n      name\n      moduleGroup {\n        moduleGroupId\n        name\n      }\n      remoteSource\n      terraformVersion\n      variables {\n        name\n        type\n        description\n        default\n      }\n      modulePropagations {\n        items {\n          name\n          modulePropagationId\n          orgUnit {\n            orgUnitId\n            name\n          }\n          orgDimension {\n            orgDimensionId\n            name\n          }\n        }\n      }\n    }\n  }\n": types.ModuleVersionDocument,
    "\n  query organizationalAccount($orgAccountId: ID!) {\n    organizationalAccount(orgAccountId: $orgAccountId) {\n      orgAccountId\n      name\n      orgUnitMemberships {\n        items {\n          orgUnit {\n            orgUnitId\n            name\n          }\n          orgDimension {\n            orgDimensionId\n            name\n          }\n        }\n      }\n      moduleAccountAssociations {\n        items {\n          status\n          modulePropagationId\n          orgAccountId\n          modulePropagation {\n            name\n            moduleGroup {\n              moduleGroupId\n              name\n            }\n            moduleVersion {\n              moduleVersionId\n              name\n            }\n          }\n        }\n      }\n    }\n  }\n": types.OrganizationalAccountDocument,
    "\n  query organizationalAccounts {\n    organizationalAccounts(limit: 100) {\n      items {\n        orgAccountId\n        name\n        cloudPlatform\n        cloudIdentifier\n      }\n    }\n  }\n": types.OrganizationalAccountsDocument,
    "\n  query organizationalDimension($orgDimensionId: ID!) {\n    organizationalDimension(orgDimensionId: $orgDimensionId) {\n      orgDimensionId\n      name\n      orgUnits {\n        items {\n          orgUnitId\n          name\n          parentOrgUnitId\n          hierarchy\n        }\n      }\n      modulePropagations {\n        items {\n          modulePropagationId\n          moduleGroupId\n          moduleVersionId\n          orgUnitId\n          orgDimensionId\n          name\n          description\n        }\n      }\n      orgUnitMemberships {\n        items {\n          orgAccount {\n            orgAccountId\n            name\n            cloudPlatform\n            cloudIdentifier\n          }\n          orgUnit {\n            orgUnitId\n            name\n          }\n        }\n      }\n    }\n  }\n": types.OrganizationalDimensionDocument,
    "\n  query organizationalDimensions {\n    organizationalDimensions(limit: 100) {\n      items {\n        orgDimensionId\n        name\n      }\n    }\n  }\n": types.OrganizationalDimensionsDocument,
    "\n  query organizationalUnit($orgUnitId: ID!, $orgDimensionId: ID!) {\n    organizationalUnit(orgDimensionId: $orgDimensionId, orgUnitId: $orgUnitId) {\n      orgUnitId\n      orgDimensionId\n      name\n      hierarchy\n      parentOrgUnitId\n      children {\n        items {\n          orgUnitId\n          name\n          hierarchy\n        }\n      }\n      downstreamOrgUnits {\n        items {\n          orgUnitId\n          name\n          hierarchy\n        }\n      }\n      orgUnitMemberships {\n        items {\n          orgAccount {\n            orgAccountId\n            name\n            cloudPlatform\n            cloudIdentifier\n          }\n        }\n      }\n      modulePropagations {\n        items {\n          modulePropagationId\n          moduleGroupId\n          moduleVersionId\n          orgUnitId\n          orgDimensionId\n          name\n          description\n        }\n      }\n    }\n  }\n": types.OrganizationalUnitDocument,
    "\n  query planExecutionRequest($planExecutionRequestId: ID!) {\n    planExecutionRequest(planExecutionRequestId: $planExecutionRequestId) {\n      planExecutionRequestId\n      status\n      requestTime\n      terraformConfigurationBase64\n      initOutput {\n        Stdout\n        Stderr\n      }\n      planOutput {\n        Stdout\n        Stderr\n      }\n    }\n  }\n": types.PlanExecutionRequestDocument,
};

/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  query applyExecutionRequest($applyExecutionRequestId: ID!) {\n    applyExecutionRequest(applyExecutionRequestId: $applyExecutionRequestId) {\n      applyExecutionRequestId\n      status\n      requestTime\n      terraformConfigurationBase64\n      initOutput {\n        Stdout\n        Stderr\n      }\n      applyOutput {\n        Stdout\n        Stderr\n      }\n    }\n  }\n"): (typeof documents)["\n  query applyExecutionRequest($applyExecutionRequestId: ID!) {\n    applyExecutionRequest(applyExecutionRequestId: $applyExecutionRequestId) {\n      applyExecutionRequestId\n      status\n      requestTime\n      terraformConfigurationBase64\n      initOutput {\n        Stdout\n        Stderr\n      }\n      applyOutput {\n        Stdout\n        Stderr\n      }\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  query moduleAccountAssociation(\n    $modulePropagationId: ID!\n    $orgAccountId: ID!\n  ) {\n    moduleAccountAssociation(\n      modulePropagationId: $modulePropagationId\n      orgAccountId: $orgAccountId\n    ) {\n      modulePropagationId\n      modulePropagation {\n        name\n        moduleGroup {\n          moduleGroupId\n          name\n        }\n        moduleVersion {\n          moduleVersionId\n          name\n        }\n      }\n      orgAccount {\n        orgAccountId\n        name\n      }\n      status\n      planExecutionRequests {\n        items {\n          planExecutionRequestId\n          status\n          requestTime\n          modulePropagationId\n          modulePropagationRequestId\n        }\n      }\n      applyExecutionRequests {\n        items {\n          applyExecutionRequestId\n          status\n          requestTime\n          modulePropagationId\n          modulePropagationRequestId\n        }\n      }\n    }\n  }\n"): (typeof documents)["\n  query moduleAccountAssociation(\n    $modulePropagationId: ID!\n    $orgAccountId: ID!\n  ) {\n    moduleAccountAssociation(\n      modulePropagationId: $modulePropagationId\n      orgAccountId: $orgAccountId\n    ) {\n      modulePropagationId\n      modulePropagation {\n        name\n        moduleGroup {\n          moduleGroupId\n          name\n        }\n        moduleVersion {\n          moduleVersionId\n          name\n        }\n      }\n      orgAccount {\n        orgAccountId\n        name\n      }\n      status\n      planExecutionRequests {\n        items {\n          planExecutionRequestId\n          status\n          requestTime\n          modulePropagationId\n          modulePropagationRequestId\n        }\n      }\n      applyExecutionRequests {\n        items {\n          applyExecutionRequestId\n          status\n          requestTime\n          modulePropagationId\n          modulePropagationRequestId\n        }\n      }\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  query moduleGroups {\n    moduleGroups(limit: 100) {\n      items {\n        moduleGroupId\n        name\n      }\n    }\n  }\n"): (typeof documents)["\n  query moduleGroups {\n    moduleGroups(limit: 100) {\n      items {\n        moduleGroupId\n        name\n      }\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  query moduleGroup($moduleGroupId: ID!) {\n    moduleGroup(moduleGroupId: $moduleGroupId) {\n      moduleGroupId\n      name\n      versions {\n        items {\n          moduleVersionId\n          name\n          remoteSource\n          terraformVersion\n        }\n      }\n      modulePropagations {\n        items {\n          name\n          modulePropagationId\n          moduleVersion {\n            moduleVersionId\n            name\n          }\n          orgUnit {\n            orgUnitId\n            name\n          }\n          orgDimension {\n            orgDimensionId\n            name\n          }\n        }\n      }\n    }\n  }\n"): (typeof documents)["\n  query moduleGroup($moduleGroupId: ID!) {\n    moduleGroup(moduleGroupId: $moduleGroupId) {\n      moduleGroupId\n      name\n      versions {\n        items {\n          moduleVersionId\n          name\n          remoteSource\n          terraformVersion\n        }\n      }\n      modulePropagations {\n        items {\n          name\n          modulePropagationId\n          moduleVersion {\n            moduleVersionId\n            name\n          }\n          orgUnit {\n            orgUnitId\n            name\n          }\n          orgDimension {\n            orgDimensionId\n            name\n          }\n        }\n      }\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  query modulePropagationExecutionRequest(\n    $modulePropagationId: ID!\n    $modulePropagationExecutionRequestId: ID!\n  ) {\n    modulePropagationExecutionRequest(\n      modulePropagationId: $modulePropagationId\n      modulePropagationExecutionRequestId: $modulePropagationExecutionRequestId\n    ) {\n      modulePropagationId\n      modulePropagationExecutionRequestId\n      requestTime\n      status\n      terraformExecutionWorkflowRequests {\n        items {\n          terraformExecutionWorkflowRequestId\n          status\n          requestTime\n          destroy\n          orgAccount {\n            orgAccountId\n            name\n          }\n          planExecutionRequest {\n            planExecutionRequestId\n            status\n            requestTime\n          }\n          applyExecutionRequest {\n            applyExecutionRequestId\n            status\n            requestTime\n          }\n        }\n      }\n      planExecutionRequests {\n        items {\n          planExecutionRequestId\n          status\n          requestTime\n          orgAccount {\n            orgAccountId\n            name\n          }\n        }\n      }\n      applyExecutionRequests {\n        items {\n          applyExecutionRequestId\n          status\n          requestTime\n          orgAccount {\n            orgAccountId\n            name\n          }\n        }\n      }\n    }\n  }\n"): (typeof documents)["\n  query modulePropagationExecutionRequest(\n    $modulePropagationId: ID!\n    $modulePropagationExecutionRequestId: ID!\n  ) {\n    modulePropagationExecutionRequest(\n      modulePropagationId: $modulePropagationId\n      modulePropagationExecutionRequestId: $modulePropagationExecutionRequestId\n    ) {\n      modulePropagationId\n      modulePropagationExecutionRequestId\n      requestTime\n      status\n      terraformExecutionWorkflowRequests {\n        items {\n          terraformExecutionWorkflowRequestId\n          status\n          requestTime\n          destroy\n          orgAccount {\n            orgAccountId\n            name\n          }\n          planExecutionRequest {\n            planExecutionRequestId\n            status\n            requestTime\n          }\n          applyExecutionRequest {\n            applyExecutionRequestId\n            status\n            requestTime\n          }\n        }\n      }\n      planExecutionRequests {\n        items {\n          planExecutionRequestId\n          status\n          requestTime\n          orgAccount {\n            orgAccountId\n            name\n          }\n        }\n      }\n      applyExecutionRequests {\n        items {\n          applyExecutionRequestId\n          status\n          requestTime\n          orgAccount {\n            orgAccountId\n            name\n          }\n        }\n      }\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  query modulePropagation($modulePropagationId: ID!) {\n    modulePropagation(modulePropagationId: $modulePropagationId) {\n      modulePropagationId\n      moduleGroup {\n        moduleGroupId\n        name\n      }\n      moduleVersion {\n        moduleVersionId\n        name\n      }\n      orgUnitId\n      orgUnit {\n        orgUnitId\n        orgDimensionId\n        name\n        downstreamOrgUnits {\n          items {\n            orgUnitId\n            orgDimensionId\n            name\n          }\n        }\n      }\n      modulePropagationId\n      name\n      description\n      executionRequests(limit: 5) {\n        items {\n          modulePropagationId\n          modulePropagationExecutionRequestId\n          requestTime\n          status\n        }\n      }\n      moduleAccountAssociations {\n        items {\n          modulePropagationId\n          orgAccount {\n            orgAccountId\n            name\n          }\n          status\n        }\n      }\n    }\n  }\n"): (typeof documents)["\n  query modulePropagation($modulePropagationId: ID!) {\n    modulePropagation(modulePropagationId: $modulePropagationId) {\n      modulePropagationId\n      moduleGroup {\n        moduleGroupId\n        name\n      }\n      moduleVersion {\n        moduleVersionId\n        name\n      }\n      orgUnitId\n      orgUnit {\n        orgUnitId\n        orgDimensionId\n        name\n        downstreamOrgUnits {\n          items {\n            orgUnitId\n            orgDimensionId\n            name\n          }\n        }\n      }\n      modulePropagationId\n      name\n      description\n      executionRequests(limit: 5) {\n        items {\n          modulePropagationId\n          modulePropagationExecutionRequestId\n          requestTime\n          status\n        }\n      }\n      moduleAccountAssociations {\n        items {\n          modulePropagationId\n          orgAccount {\n            orgAccountId\n            name\n          }\n          status\n        }\n      }\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  mutation createModulePropagationExecutionRequest($modulePropagationId: ID!) {\n    createModulePropagationExecutionRequest(\n      modulePropagationExecutionRequest: {\n        modulePropagationId: $modulePropagationId\n      }\n    ) {\n      modulePropagationExecutionRequestId\n      status\n    }\n  }\n"): (typeof documents)["\n  mutation createModulePropagationExecutionRequest($modulePropagationId: ID!) {\n    createModulePropagationExecutionRequest(\n      modulePropagationExecutionRequest: {\n        modulePropagationId: $modulePropagationId\n      }\n    ) {\n      modulePropagationExecutionRequestId\n      status\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  query orgDimensions {\n    organizationalDimensions(limit: 10000) {\n      items {\n        orgDimensionId\n        name\n        orgUnits(limit: 10000) {\n          items {\n            orgUnitId\n            name\n          }\n        }\n      }\n    }\n  }\n"): (typeof documents)["\n  query orgDimensions {\n    organizationalDimensions(limit: 10000) {\n      items {\n        orgDimensionId\n        name\n        orgUnits(limit: 10000) {\n          items {\n            orgUnitId\n            name\n          }\n        }\n      }\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  mutation updateModulePropagation(\n    $modulePropagationId: ID!\n    $update: ModulePropagationUpdate!\n  ) {\n    updateModulePropagation(\n      modulePropagationId: $modulePropagationId\n      update: $update\n    ) {\n      modulePropagationId\n    }\n  }\n"): (typeof documents)["\n  mutation updateModulePropagation(\n    $modulePropagationId: ID!\n    $update: ModulePropagationUpdate!\n  ) {\n    updateModulePropagation(\n      modulePropagationId: $modulePropagationId\n      update: $update\n    ) {\n      modulePropagationId\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  query moduleVersion($moduleGroupId: ID!, $moduleVersionId: ID!) {\n    moduleVersion(\n      moduleGroupId: $moduleGroupId\n      moduleVersionId: $moduleVersionId\n    ) {\n      moduleVersionId\n      name\n      moduleGroup {\n        moduleGroupId\n        name\n      }\n      remoteSource\n      terraformVersion\n      variables {\n        name\n        type\n        description\n        default\n      }\n      modulePropagations {\n        items {\n          name\n          modulePropagationId\n          orgUnit {\n            orgUnitId\n            name\n          }\n          orgDimension {\n            orgDimensionId\n            name\n          }\n        }\n      }\n    }\n  }\n"): (typeof documents)["\n  query moduleVersion($moduleGroupId: ID!, $moduleVersionId: ID!) {\n    moduleVersion(\n      moduleGroupId: $moduleGroupId\n      moduleVersionId: $moduleVersionId\n    ) {\n      moduleVersionId\n      name\n      moduleGroup {\n        moduleGroupId\n        name\n      }\n      remoteSource\n      terraformVersion\n      variables {\n        name\n        type\n        description\n        default\n      }\n      modulePropagations {\n        items {\n          name\n          modulePropagationId\n          orgUnit {\n            orgUnitId\n            name\n          }\n          orgDimension {\n            orgDimensionId\n            name\n          }\n        }\n      }\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  query organizationalAccount($orgAccountId: ID!) {\n    organizationalAccount(orgAccountId: $orgAccountId) {\n      orgAccountId\n      name\n      orgUnitMemberships {\n        items {\n          orgUnit {\n            orgUnitId\n            name\n          }\n          orgDimension {\n            orgDimensionId\n            name\n          }\n        }\n      }\n      moduleAccountAssociations {\n        items {\n          status\n          modulePropagationId\n          orgAccountId\n          modulePropagation {\n            name\n            moduleGroup {\n              moduleGroupId\n              name\n            }\n            moduleVersion {\n              moduleVersionId\n              name\n            }\n          }\n        }\n      }\n    }\n  }\n"): (typeof documents)["\n  query organizationalAccount($orgAccountId: ID!) {\n    organizationalAccount(orgAccountId: $orgAccountId) {\n      orgAccountId\n      name\n      orgUnitMemberships {\n        items {\n          orgUnit {\n            orgUnitId\n            name\n          }\n          orgDimension {\n            orgDimensionId\n            name\n          }\n        }\n      }\n      moduleAccountAssociations {\n        items {\n          status\n          modulePropagationId\n          orgAccountId\n          modulePropagation {\n            name\n            moduleGroup {\n              moduleGroupId\n              name\n            }\n            moduleVersion {\n              moduleVersionId\n              name\n            }\n          }\n        }\n      }\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  query organizationalAccounts {\n    organizationalAccounts(limit: 100) {\n      items {\n        orgAccountId\n        name\n        cloudPlatform\n        cloudIdentifier\n      }\n    }\n  }\n"): (typeof documents)["\n  query organizationalAccounts {\n    organizationalAccounts(limit: 100) {\n      items {\n        orgAccountId\n        name\n        cloudPlatform\n        cloudIdentifier\n      }\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  query organizationalDimension($orgDimensionId: ID!) {\n    organizationalDimension(orgDimensionId: $orgDimensionId) {\n      orgDimensionId\n      name\n      orgUnits {\n        items {\n          orgUnitId\n          name\n          parentOrgUnitId\n          hierarchy\n        }\n      }\n      modulePropagations {\n        items {\n          modulePropagationId\n          moduleGroupId\n          moduleVersionId\n          orgUnitId\n          orgDimensionId\n          name\n          description\n        }\n      }\n      orgUnitMemberships {\n        items {\n          orgAccount {\n            orgAccountId\n            name\n            cloudPlatform\n            cloudIdentifier\n          }\n          orgUnit {\n            orgUnitId\n            name\n          }\n        }\n      }\n    }\n  }\n"): (typeof documents)["\n  query organizationalDimension($orgDimensionId: ID!) {\n    organizationalDimension(orgDimensionId: $orgDimensionId) {\n      orgDimensionId\n      name\n      orgUnits {\n        items {\n          orgUnitId\n          name\n          parentOrgUnitId\n          hierarchy\n        }\n      }\n      modulePropagations {\n        items {\n          modulePropagationId\n          moduleGroupId\n          moduleVersionId\n          orgUnitId\n          orgDimensionId\n          name\n          description\n        }\n      }\n      orgUnitMemberships {\n        items {\n          orgAccount {\n            orgAccountId\n            name\n            cloudPlatform\n            cloudIdentifier\n          }\n          orgUnit {\n            orgUnitId\n            name\n          }\n        }\n      }\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  query organizationalDimensions {\n    organizationalDimensions(limit: 100) {\n      items {\n        orgDimensionId\n        name\n      }\n    }\n  }\n"): (typeof documents)["\n  query organizationalDimensions {\n    organizationalDimensions(limit: 100) {\n      items {\n        orgDimensionId\n        name\n      }\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  query organizationalUnit($orgUnitId: ID!, $orgDimensionId: ID!) {\n    organizationalUnit(orgDimensionId: $orgDimensionId, orgUnitId: $orgUnitId) {\n      orgUnitId\n      orgDimensionId\n      name\n      hierarchy\n      parentOrgUnitId\n      children {\n        items {\n          orgUnitId\n          name\n          hierarchy\n        }\n      }\n      downstreamOrgUnits {\n        items {\n          orgUnitId\n          name\n          hierarchy\n        }\n      }\n      orgUnitMemberships {\n        items {\n          orgAccount {\n            orgAccountId\n            name\n            cloudPlatform\n            cloudIdentifier\n          }\n        }\n      }\n      modulePropagations {\n        items {\n          modulePropagationId\n          moduleGroupId\n          moduleVersionId\n          orgUnitId\n          orgDimensionId\n          name\n          description\n        }\n      }\n    }\n  }\n"): (typeof documents)["\n  query organizationalUnit($orgUnitId: ID!, $orgDimensionId: ID!) {\n    organizationalUnit(orgDimensionId: $orgDimensionId, orgUnitId: $orgUnitId) {\n      orgUnitId\n      orgDimensionId\n      name\n      hierarchy\n      parentOrgUnitId\n      children {\n        items {\n          orgUnitId\n          name\n          hierarchy\n        }\n      }\n      downstreamOrgUnits {\n        items {\n          orgUnitId\n          name\n          hierarchy\n        }\n      }\n      orgUnitMemberships {\n        items {\n          orgAccount {\n            orgAccountId\n            name\n            cloudPlatform\n            cloudIdentifier\n          }\n        }\n      }\n      modulePropagations {\n        items {\n          modulePropagationId\n          moduleGroupId\n          moduleVersionId\n          orgUnitId\n          orgDimensionId\n          name\n          description\n        }\n      }\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  query planExecutionRequest($planExecutionRequestId: ID!) {\n    planExecutionRequest(planExecutionRequestId: $planExecutionRequestId) {\n      planExecutionRequestId\n      status\n      requestTime\n      terraformConfigurationBase64\n      initOutput {\n        Stdout\n        Stderr\n      }\n      planOutput {\n        Stdout\n        Stderr\n      }\n    }\n  }\n"): (typeof documents)["\n  query planExecutionRequest($planExecutionRequestId: ID!) {\n    planExecutionRequest(planExecutionRequestId: $planExecutionRequestId) {\n      planExecutionRequestId\n      status\n      requestTime\n      terraformConfigurationBase64\n      initOutput {\n        Stdout\n        Stderr\n      }\n      planOutput {\n        Stdout\n        Stderr\n      }\n    }\n  }\n"];

/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 *
 *
 * @example
 * ```ts
 * const query = gql(`query GetUser($id: ID!) { user(id: $id) { name } }`);
 * ```
 *
 * The query argument is unknown!
 * Please regenerate the types.
**/
export function gql(source: string): unknown;

export function gql(source: string) {
  return (documents as any)[source] ?? {};
}

export type DocumentType<TDocumentNode extends DocumentNode<any, any>> = TDocumentNode extends DocumentNode<  infer TType,  any>  ? TType  : never;