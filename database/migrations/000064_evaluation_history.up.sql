-- SPDX-FileCopyrightText: Copyright 2024 The Minder Authors
-- SPDX-License-Identifier: Apache-2.0

BEGIN;

CREATE TABLE rule_entity_evaluations(
    id           UUID NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
    rule_id      UUID NOT NULL REFERENCES rule_instances(id) ON DELETE CASCADE,
    -- Copying the same pattern of linking to entity types as used in rule_evaluations
    repository_id UUID REFERENCES repositories(id),
    pull_request_id UUID REFERENCES pull_requests(id),
    artifact_id UUID REFERENCES artifacts(id),
    UNIQUE(rule_id, repository_id),
    UNIQUE(rule_id, pull_request_id),
    UNIQUE(rule_id, artifact_id),
    -- exactly one entity ID column must be set
    CONSTRAINT one_entity_id CHECK (num_nonnulls(repository_id, artifact_id, pull_request_id) = 1)
);

CREATE TABLE evaluation_history(
    id UUID NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
    rule_entity_id UUID NOT NULL REFERENCES rule_entity_evaluations(id) ON DELETE CASCADE,
    status eval_status_types NOT NULL,
    details TEXT NOT NULL
);

CREATE TABLE evaluation_instance(
    evaluation_id UUID NOT NULL REFERENCES evaluation_history(id) ON DELETE CASCADE,
    evaluation_time TIMESTAMP NOT NULL DEFAULT NOW(),
    PRIMARY KEY (evaluation_id, evaluation_time)
);

CREATE TABLE latest_evaluation_state(
    rule_entity_id UUID NOT NULL PRIMARY KEY REFERENCES rule_entity_evaluations(id) ON DELETE CASCADE,
    evaluation_history_id UUID NOT NULL REFERENCES evaluation_history(id)
);

CREATE TABLE remediation_events(
    id UUID NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
    evaluation_id UUID NOT NULL REFERENCES evaluation_history(id),
    status remediation_status_types NOT NULL,
    details TEXT NOT NULL,
    metadata JSONB NOT NULL DEFAULT '{}'::JSONB,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE alert_events(
    id UUID NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
    evaluation_id UUID NOT NULL REFERENCES evaluation_history(id),
    status alert_status_types NOT NULL,
    details TEXT NOT NULL,
    metadata JSONB NOT NULL DEFAULT '{}'::JSONB,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

COMMIT;
