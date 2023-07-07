-- CreateTable
CREATE TABLE "delegators" (
    "address" VARCHAR(36) NOT NULL,
    "first_seen" TIMESTAMP(3),
    "alias" TEXT,

    CONSTRAINT "delegators_pkey" PRIMARY KEY ("address")
);

-- Seed insert from delegations if any
INSERT INTO "delegators" ("address")
SELECT DISTINCT delegator
FROM "delegations";
-- -- Seed first_seen
UPDATE "delegators"
SET first_seen =(
    SELECT  "delegations".timestamp
    FROM "delegations"
    WHERE  "delegators".address = "delegations".delegator
    ORDER BY timestamp limit 1
);

-- AddForeignKey
ALTER TABLE "delegations" ADD CONSTRAINT "delegations_delegator_fkey" FOREIGN KEY ("delegator") REFERENCES "delegators"("address") ON DELETE RESTRICT ON UPDATE CASCADE;
