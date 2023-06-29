-- CreateTable
CREATE TABLE "delegations" (
    "id" BIGINT NOT NULL,
    "timestamp" TIMESTAMP(3) NOT NULL,
    "amount" BIGINT NOT NULL,
    "delegator" VARCHAR(36) NOT NULL,
    "block_hash" TEXT NOT NULL,
    "block_level" BIGINT NOT NULL,

    CONSTRAINT "delegations_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE INDEX "delegations_delegator_idx" ON "delegations"("delegator");

-- CreateIndex
CREATE INDEX "delegations_block_hash_idx" ON "delegations"("block_hash");

-- CreateIndex
CREATE INDEX "delegations_block_level_idx" ON "delegations"("block_level");

-- CreateIndex
CREATE INDEX "delegations_timestamp_idx" ON "delegations"("timestamp");
