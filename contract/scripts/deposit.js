const hre = require("hardhat");

async function main() {
  const [owner, user1] = await hre.ethers.getSigners();
  
  console.log("Owner address:", owner.address);
  console.log("User1 address:", user1.address);
  
  const AuctionVault = await hre.ethers.getContractAt(
    "AuctionVault", 
    "0x5FbDB2315678afecb367f032d93F642f64180aa3"
  );

  // Deposit 1 ETH from user1
  console.log("Making deposit of 1 ETH...");
  const tx = await AuctionVault.connect(user1).deposit({
    value: hre.ethers.parseEther("1.0")
  });
  
  await tx.wait();
  console.log("✅ Deposited 1 ETH from:", user1.address);
  
  // Check balance on contract
  const balance = await AuctionVault.balances(user1.address);
  console.log("Contract balance for user:", hre.ethers.formatEther(balance), "ETH");
}

main().catch(console.error);
