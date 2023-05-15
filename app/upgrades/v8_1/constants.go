// Copyright 2022 Gridiron Foundation
// This file is part of the Gridiron Network packages.
//
// Gridiron is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Gridiron packages are distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Gridiron packages. If not, see https://github.com/gridchain/gridiron/blob/main/LICENSE

package v81

const (
	// UpgradeName is the shared upgrade plan name for mainnet and testnet
	UpgradeName = "v8.1.0"
	// MainnetUpgradeHeight defines the Gridiron mainnet block height on which the upgrade will take place
	MainnetUpgradeHeight = 4_500_000
	// TestnetUpgradeHeight defines the Gridiron testnet block height on which the upgrade will take place
	TestnetUpgradeHeight = 5_280_000
	// UpgradeInfo defines the binaries that will be used for the upgrade
	UpgradeInfo = `'{"binaries":{"darwin/arm64":"https://github.com/gridchain/gridiron/releases/download/v8.1.0/gridiron_8.1.0_Darwin_arm64.tar.gz","darwin/x86_64":"https://github.com/gridchain/gridiron/releases/download/v8.1.0/gridiron_8.1.0_Darwin_x86_64.tar.gz","linux/arm64":"https://github.com/gridchain/gridiron/releases/download/v8.1.0/gridiron_8.1.0_Linux_arm64.tar.gz","linux/amd64":"https://github.com/gridchain/gridiron/releases/download/v8.1.0/gridiron_8.1.0_Linux_x86_64.tar.gz","windows/x86_64":"https://github.com/gridchain/gridiron/releases/download/v8.1.0/gridiron_8.1.0_Windows_x86_64.zip"}}'`
)
