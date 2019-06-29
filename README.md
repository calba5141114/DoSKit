# DoSKit
In computing, a denial-of-service attack (DoS attack) is a cyber-attack in which the perpetrator seeks to make a machine or network resource unavailable to its intended users by temporarily or indefinitely disrupting services of a host connected to the Internet. Denial of service is typically accomplished by flooding the targeted machine or resource with superfluous requests in an attempt to overload systems and prevent some or all legitimate requests from being fulfilled.

# Purpose 
DoSKit is a security testing tool for Denial of Service attacks 
directed at one domain or another. This tool is intended to be used 
for stress testing and simulated DoS Attack. Never run DoSKit on a 
domain without the explicit permission of the owner. 

# Goals 
- Send 1 million requests per second with little memory usage. 
- Prevent useless resource usage. 
- Store Targets in a SQLITE DB that can be pulled up via the CLI
- Colors on the terminal (rgb)
- DoSKit Microservice with REST (far out to potential liability and abuse)

# News 
### CLI 
CLI is working and can successfully execute a DoS Attack.
Beware if you play with this CLI on servers that are not your own you can suffer grave consequences and your IP can be permanently blocked from 
attacked sites.

### DDoS
CLI can be run on any number of hosts for a DDoS attack 
utilizing tools like Salt with compiled binaries or by compiling 
the Docker File in the project and using Scheduling tools

### Running Compiled Binary on the CLI
./DosKit --target="https://www.random.com"
