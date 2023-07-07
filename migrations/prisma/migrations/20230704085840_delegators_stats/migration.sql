-- CreateEnum
CREATE TYPE "stat_kind" AS ENUM ('TOP10VALIDATORS', 'TOP100VALIDATORS', 'DAILYVOLUME', 'WEEKLYVOLUME', 'MONTHLYVOLUME', 'YEARLYVOLUME');

-- CreateTable
CREATE TABLE "delegations_stats" (
    "id" BIGSERIAL NOT NULL,
    "timestamp" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "kind" "stat_kind" NOT NULL,
    "value" JSONB NOT NULL,
    "delegator_address" VARCHAR(36),

    CONSTRAINT "delegations_stats_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE INDEX "delegations_stats_timestamp_idx" ON "delegations_stats"("timestamp");

-- CreateIndex
CREATE INDEX "delegations_stats_kind_idx" ON "delegations_stats"("kind");

-- AddForeignKey
ALTER TABLE "delegations_stats" ADD CONSTRAINT "delegations_stats_delegator_address_fkey" FOREIGN KEY ("delegator_address") REFERENCES "delegators"("address") ON DELETE SET NULL ON UPDATE CASCADE;
