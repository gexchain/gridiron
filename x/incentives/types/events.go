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
// along with the Gridiron packages. If not, see https://github.com/gridiron/gridiron/blob/main/LICENSE

package types

// incentives events
const (
	EventTypeRegisterIncentive    = "register_incentive"
	EventTypeCancelIncentive      = "cancel_incentive"
	EventTypeDistributeIncentives = "distribute_incentives"

	AttributeKeyContract = "contract"
	AttributeKeyEpochs   = "epochs"
)
