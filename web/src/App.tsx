import React from "react";
import { Routes, Route } from "react-router-dom";

import { OrganizationalDimensionsList } from "./components/OrganizationalDimensionsList";
import { ModulePropagationExecutionRequestPage } from "./components/ModulePropagationExecutionPage";
import { ModulePropagationPage } from "./components/ModulePropagationPage";
import { PageWrapper } from "./components/PageWrapper";
import { PlanExecutionRequestPage } from "./components/PlanExecutionRequestPage";
import { ApplyExecutionRequestPage } from "./components/ApplyExecutionRequestPage";
import { OrganizationalAccountsList } from "./components/OrganizationalAccountsList";
import { OrganizationalAccountPage } from "./components/OrganizationalAccountPage";
import { OrganizationalDimensionPage } from "./components/OrganizationalDimensionPage";
import { OrganizationalUnitPage } from "./components/OrganizationalUnitPage";
import { ModuleAccountAssociationPage } from "./components/ModuleAccountAssociationPage";

function App() {
  return (
    <PageWrapper>
      <Routes>
        <Route path="/" element={<OrganizationalDimensionsList />} />
        <Route
          path="/module-propagations/:modulePropagationId"
          element={<ModulePropagationPage />}
        />
        <Route
          path="/module-propagations/:modulePropagationId/executions/:modulePropagationExecutionRequestId"
          element={<ModulePropagationExecutionRequestPage />}
        />
        <Route
          path="/module-propagations/:modulePropagationId/account-associations/:orgAccountId"
          element={<ModuleAccountAssociationPage />}
        />
        <Route
          path="/plan-execution-requests/:planExecutionRequestId"
          element={<PlanExecutionRequestPage />}
        />
        <Route
          path="/apply-execution-requests/:applyExecutionRequestId"
          element={<ApplyExecutionRequestPage />}
        />
        <Route
          path="/org-dimensions"
          element={<OrganizationalDimensionsList />}
        />
        <Route
          path="/org-dimensions/:organizationalDimensionId"
          element={<OrganizationalDimensionPage />}
        />
        <Route
          path="/org-dimensions/:organizationalDimensionId/org-units/:organizationalUnitId"
          element={<OrganizationalUnitPage />}
        />
        <Route path="/org-accounts" element={<OrganizationalAccountsList />} />
        <Route
          path="/org-accounts/:organizationalAccountId"
          element={<OrganizationalAccountPage />}
        />
      </Routes>
    </PageWrapper>
  );
}

export default App;
