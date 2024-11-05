# Hyperledger Fabric

## Overview
Hyperledger Fabric is a Graduated project under the Hyperledger umbrella, designed for distributed ledger solutions. Its modular architecture provides high levels of confidentiality, resiliency, flexibility, and scalability. Hyperledger Fabric allows for pluggable implementations of various components, accommodating the complexities of different economic ecosystems. This platform offers a uniquely elastic and extensible architecture, setting it apart from other blockchain solutions. Building on a fully-vetted, open-source framework, Hyperledger Fabric is an ideal starting point for enterprise blockchain initiatives.

## Releases
Hyperledger Fabric provides periodic releases with new features and improvements. Certain releases are designated as Long-Term Support (LTS), ensuring that important fixes are backported during overlap periods.

### Current LTS Release:
- v2.5.x

### Historic LTS Releases:
- v2.2.x (maintenance ended February 2024)
- v1.4.x (maintenance ended April 2021)

For complete release notes, visit the [GitHub releases page](https://github.com/hyperledger/fabric/releases).

## Documentation and Getting Started
To familiarize yourself with Hyperledger Fabric, visit our comprehensive online documentation:
- [Getting Started with v2.5](https://hyperledger-fabric.readthedocs.io/en/release-2.5/getting-started.html)
- [Previous Versions](https://hyperledger-fabric.readthedocs.io/en/release-2.5/previous_versions.html)

We recommend that first-time users start with the Getting Started section to understand the components and basic transaction flow.

## Contributing
We welcome contributions to Hyperledger Fabric in various forms. There’s always plenty to do! Check our [contribution guidelines](https://hyperledger-fabric.readthedocs.io/en/release-2.5/contributing.html) for more details on how to get involved.

## Community
Engage with the Hyperledger community:
- [Hyperledger Community Meetup](https://www.hyperledger.org/community/community-meetup)
- [Mailing Lists and Archives](https://lists.hyperledger.org/g/hyperledger-fabric)
- [Discord Chat](https://chat.hyperledger.org/)
- [Issue Tracking](https://jira.hyperledger.org/projects/FABRIC)

## License
Hyperledger Fabric source code is available under the Apache License, Version 2.0 (Apache-2.0), and documentation files are under the Creative Commons Attribution 4.0 International License (CC-BY-4.0).

---

# SmartFalcon Internship - Hyperledger Fabric Assignment

## Overview
This project is developed as part of the SmartFalcon internship program. The objective is to create a blockchain-based system using Hyperledger Fabric to manage and track financial assets. The system supports asset creation, updating asset values, querying the world state, and retrieving transaction history, ensuring security, transparency, and immutability of asset records.

## Table of Contents
- [Project Structure](#project-structure)
- [Technologies Used](#technologies-used)
- [Installation Instructions](#installation-instructions)
- [Setup Instructions](#setup-instructions)
- [Smart Contract](#smart-contract)
- [REST API](#rest-api)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Project Structure

SmartFalconInternship/ 
├── docker-compose.yml # Docker configuration for the Hyperledger Fabric network 
├── rest-api/ # REST API project folder  
| ├── Dockerfile # Dockerfile for containerizing the REST API 
│ ├── app.js # Main application file 
| ├── package.json # NPM package configuration 
│ └── routes/ # Folder for API routes 
│ └── assetRoutes.js # API routes for asset management 
└── smart-contract/ # Smart contract folder 
├── assetContract.go # Smart contract code for asset management 
└── README.md # Documentation for smart contract

## Technologies Used
- **Hyperledger Fabric**: A permissioned blockchain framework.
- **Golang**: Programming language used for developing the smart contract.
- **Node.js**: Environment for running the REST API.
- **Docker**: Containerization platform for running the application.
- **GitHub**: Version control for source code management.

## Installation Instructions
1. **Clone the repository**:
    ```bash
    git clone https://github.com/buntyanvi/SmartFalconInternship.git
    cd SmartFalconInternship
    ```

2. **Install dependencies for the REST API**:
    ```bash
    cd rest-api
    npm install
    ```

## Hyperledger Fabric Test Network Setup
### Ensure Docker is Installed and Running
Make sure Docker is installed and running on your machine.

## Setup Instructions
### Hyperledger Fabric Test Network
1. Navigate to the Hyperledger Fabric test network directory:
    ```bash
    cd hyperledger-fabric-test-network
    ```

2. Start the test network:
    ```bash
    ./network.sh up
    ```

3. Deploy the smart contract:
    ```bash
    ./network.sh deployCC
    ```

### REST API
1. Build the Docker image for the REST API:
    ```bash
    docker build -t rest-api .
    ```

2. Run the Docker container:
    ```bash
    docker run -p 3000:3000 rest-api
    ```

## Smart Contract
The smart contract is implemented in **Golang** and located in the `smart-contract` folder. The contract handles the creation, update, and querying of assets with attributes such as:
- **DEALERID**
- **MSISDN**
- **MPIN**
- **BALANCE**
- **STATUS**
- **TRANSAMOUNT**
- **TRANSTYPE**
- **REMARKS**

### Key Functions
- **CreateAsset**: Creates a new asset.
- **UpdateAsset**: Updates the value of an existing asset.
- **QueryAsset**: Retrieves asset details.
- **GetTransactionHistory**: Returns the transaction history for an asset.

## REST API
The REST API is implemented using **Node.js** and provides endpoints for interacting with the smart contract.

### Endpoints
- **POST** `/assets`: Create a new asset.
- **PUT** `/assets/{id}`: Update an existing asset.
- **GET** `/assets/{id}`: Retrieve details of an asset.
- **GET** `/assets/history/{id}`: Get transaction history for an asset.

## Usage
Use Postman or any API testing tool to interact with the REST API endpoints. Make sure the Hyperledger Fabric test network is up and running before making API calls.

## Contributing
Contributions are welcome! Please fork the repository and create a pull request.

## License
This project is licensed under the MIT License. See the `LICENSE` file for more information.
