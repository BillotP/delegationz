-- CreateTable
CREATE TABLE "delegations" (
    "id" BIGINT NOT NULL,
    "timestamp" TIMESTAMP(3) NOT NULL,
    "amount" BIGINT NOT NULL,
    "delegator" VARCHAR(36) NOT NULL,
    "block" BIGINT NOT NULL,

    CONSTRAINT "delegations_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE INDEX "delegations_delegator_idx" ON "delegations"("delegator");

-- CreateIndex
CREATE INDEX "delegations_block_idx" ON "delegations"("block");

-- CreateIndex
CREATE INDEX "delegations_timestamp_idx" ON "delegations"("timestamp");
