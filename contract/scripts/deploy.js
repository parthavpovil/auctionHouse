const hre = require("hardhat");

async function main() {
  const AuctionVault = await hre.ethers.getContractFactory("AuctionVault");
  const auctionVault = await AuctionVault.deploy();

  await auctionVault.waitForDeployment();

  console.log("AuctionVault deployed to:", await auctionVault.getAddress());
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
