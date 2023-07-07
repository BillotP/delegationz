/*
  Warnings:

  - A unique constraint covering the columns `[value]` on the table `delegations_stats` will be added. If there are existing duplicate values, this will fail.
  - A unique constraint covering the columns `[kind,timestamp]` on the table `delegations_stats` will be added. If there are existing duplicate values, this will fail.

*/
-- CreateIndex
CREATE UNIQUE INDEX "delegations_stats_value_key" ON "delegations_stats"("value");

-- CreateIndex
CREATE UNIQUE INDEX "delegations_stats_kind_timestamp_key" ON "delegations_stats"("kind", "timestamp");
