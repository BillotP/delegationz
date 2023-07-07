/*
  Warnings:

  - Made the column `first_seen` on table `delegators` required. This step will fail if there are existing NULL values in that column.

*/
-- AlterTable
ALTER TABLE "delegators" ALTER COLUMN "first_seen" SET NOT NULL;
