-- CreateTable
CREATE TABLE "url" (
    "id" TEXT NOT NULL,
    "shortCode" TEXT NOT NULL,
    "originalURL" TEXT NOT NULL,
    "clicks" INTEGER NOT NULL DEFAULT 0,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT "url_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE UNIQUE INDEX "url_shortCode_key" ON "url"("shortCode");
