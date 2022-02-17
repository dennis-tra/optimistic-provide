// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("CloserPeers", testCloserPeers)
	t.Run("Connections", testConnections)
	t.Run("Dials", testDials)
	t.Run("FindNodes", testFindNodes)
	t.Run("IPAddresses", testIPAddresses)
	t.Run("MultiAddresses", testMultiAddresses)
	t.Run("PeerLogs", testPeerLogs)
	t.Run("PeerStates", testPeerStates)
	t.Run("Peers", testPeers)
	t.Run("Provides", testProvides)
	t.Run("RoutingTableEntries", testRoutingTableEntries)
	t.Run("RoutingTableSnapshots", testRoutingTableSnapshots)
}

func TestDelete(t *testing.T) {
	t.Run("CloserPeers", testCloserPeersDelete)
	t.Run("Connections", testConnectionsDelete)
	t.Run("Dials", testDialsDelete)
	t.Run("FindNodes", testFindNodesDelete)
	t.Run("IPAddresses", testIPAddressesDelete)
	t.Run("MultiAddresses", testMultiAddressesDelete)
	t.Run("PeerLogs", testPeerLogsDelete)
	t.Run("PeerStates", testPeerStatesDelete)
	t.Run("Peers", testPeersDelete)
	t.Run("Provides", testProvidesDelete)
	t.Run("RoutingTableEntries", testRoutingTableEntriesDelete)
	t.Run("RoutingTableSnapshots", testRoutingTableSnapshotsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("CloserPeers", testCloserPeersQueryDeleteAll)
	t.Run("Connections", testConnectionsQueryDeleteAll)
	t.Run("Dials", testDialsQueryDeleteAll)
	t.Run("FindNodes", testFindNodesQueryDeleteAll)
	t.Run("IPAddresses", testIPAddressesQueryDeleteAll)
	t.Run("MultiAddresses", testMultiAddressesQueryDeleteAll)
	t.Run("PeerLogs", testPeerLogsQueryDeleteAll)
	t.Run("PeerStates", testPeerStatesQueryDeleteAll)
	t.Run("Peers", testPeersQueryDeleteAll)
	t.Run("Provides", testProvidesQueryDeleteAll)
	t.Run("RoutingTableEntries", testRoutingTableEntriesQueryDeleteAll)
	t.Run("RoutingTableSnapshots", testRoutingTableSnapshotsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("CloserPeers", testCloserPeersSliceDeleteAll)
	t.Run("Connections", testConnectionsSliceDeleteAll)
	t.Run("Dials", testDialsSliceDeleteAll)
	t.Run("FindNodes", testFindNodesSliceDeleteAll)
	t.Run("IPAddresses", testIPAddressesSliceDeleteAll)
	t.Run("MultiAddresses", testMultiAddressesSliceDeleteAll)
	t.Run("PeerLogs", testPeerLogsSliceDeleteAll)
	t.Run("PeerStates", testPeerStatesSliceDeleteAll)
	t.Run("Peers", testPeersSliceDeleteAll)
	t.Run("Provides", testProvidesSliceDeleteAll)
	t.Run("RoutingTableEntries", testRoutingTableEntriesSliceDeleteAll)
	t.Run("RoutingTableSnapshots", testRoutingTableSnapshotsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("CloserPeers", testCloserPeersExists)
	t.Run("Connections", testConnectionsExists)
	t.Run("Dials", testDialsExists)
	t.Run("FindNodes", testFindNodesExists)
	t.Run("IPAddresses", testIPAddressesExists)
	t.Run("MultiAddresses", testMultiAddressesExists)
	t.Run("PeerLogs", testPeerLogsExists)
	t.Run("PeerStates", testPeerStatesExists)
	t.Run("Peers", testPeersExists)
	t.Run("Provides", testProvidesExists)
	t.Run("RoutingTableEntries", testRoutingTableEntriesExists)
	t.Run("RoutingTableSnapshots", testRoutingTableSnapshotsExists)
}

func TestFind(t *testing.T) {
	t.Run("CloserPeers", testCloserPeersFind)
	t.Run("Connections", testConnectionsFind)
	t.Run("Dials", testDialsFind)
	t.Run("FindNodes", testFindNodesFind)
	t.Run("IPAddresses", testIPAddressesFind)
	t.Run("MultiAddresses", testMultiAddressesFind)
	t.Run("PeerLogs", testPeerLogsFind)
	t.Run("PeerStates", testPeerStatesFind)
	t.Run("Peers", testPeersFind)
	t.Run("Provides", testProvidesFind)
	t.Run("RoutingTableEntries", testRoutingTableEntriesFind)
	t.Run("RoutingTableSnapshots", testRoutingTableSnapshotsFind)
}

func TestBind(t *testing.T) {
	t.Run("CloserPeers", testCloserPeersBind)
	t.Run("Connections", testConnectionsBind)
	t.Run("Dials", testDialsBind)
	t.Run("FindNodes", testFindNodesBind)
	t.Run("IPAddresses", testIPAddressesBind)
	t.Run("MultiAddresses", testMultiAddressesBind)
	t.Run("PeerLogs", testPeerLogsBind)
	t.Run("PeerStates", testPeerStatesBind)
	t.Run("Peers", testPeersBind)
	t.Run("Provides", testProvidesBind)
	t.Run("RoutingTableEntries", testRoutingTableEntriesBind)
	t.Run("RoutingTableSnapshots", testRoutingTableSnapshotsBind)
}

func TestOne(t *testing.T) {
	t.Run("CloserPeers", testCloserPeersOne)
	t.Run("Connections", testConnectionsOne)
	t.Run("Dials", testDialsOne)
	t.Run("FindNodes", testFindNodesOne)
	t.Run("IPAddresses", testIPAddressesOne)
	t.Run("MultiAddresses", testMultiAddressesOne)
	t.Run("PeerLogs", testPeerLogsOne)
	t.Run("PeerStates", testPeerStatesOne)
	t.Run("Peers", testPeersOne)
	t.Run("Provides", testProvidesOne)
	t.Run("RoutingTableEntries", testRoutingTableEntriesOne)
	t.Run("RoutingTableSnapshots", testRoutingTableSnapshotsOne)
}

func TestAll(t *testing.T) {
	t.Run("CloserPeers", testCloserPeersAll)
	t.Run("Connections", testConnectionsAll)
	t.Run("Dials", testDialsAll)
	t.Run("FindNodes", testFindNodesAll)
	t.Run("IPAddresses", testIPAddressesAll)
	t.Run("MultiAddresses", testMultiAddressesAll)
	t.Run("PeerLogs", testPeerLogsAll)
	t.Run("PeerStates", testPeerStatesAll)
	t.Run("Peers", testPeersAll)
	t.Run("Provides", testProvidesAll)
	t.Run("RoutingTableEntries", testRoutingTableEntriesAll)
	t.Run("RoutingTableSnapshots", testRoutingTableSnapshotsAll)
}

func TestCount(t *testing.T) {
	t.Run("CloserPeers", testCloserPeersCount)
	t.Run("Connections", testConnectionsCount)
	t.Run("Dials", testDialsCount)
	t.Run("FindNodes", testFindNodesCount)
	t.Run("IPAddresses", testIPAddressesCount)
	t.Run("MultiAddresses", testMultiAddressesCount)
	t.Run("PeerLogs", testPeerLogsCount)
	t.Run("PeerStates", testPeerStatesCount)
	t.Run("Peers", testPeersCount)
	t.Run("Provides", testProvidesCount)
	t.Run("RoutingTableEntries", testRoutingTableEntriesCount)
	t.Run("RoutingTableSnapshots", testRoutingTableSnapshotsCount)
}

func TestHooks(t *testing.T) {
	t.Run("CloserPeers", testCloserPeersHooks)
	t.Run("Connections", testConnectionsHooks)
	t.Run("Dials", testDialsHooks)
	t.Run("FindNodes", testFindNodesHooks)
	t.Run("IPAddresses", testIPAddressesHooks)
	t.Run("MultiAddresses", testMultiAddressesHooks)
	t.Run("PeerLogs", testPeerLogsHooks)
	t.Run("PeerStates", testPeerStatesHooks)
	t.Run("Peers", testPeersHooks)
	t.Run("Provides", testProvidesHooks)
	t.Run("RoutingTableEntries", testRoutingTableEntriesHooks)
	t.Run("RoutingTableSnapshots", testRoutingTableSnapshotsHooks)
}

func TestInsert(t *testing.T) {
	t.Run("CloserPeers", testCloserPeersInsert)
	t.Run("CloserPeers", testCloserPeersInsertWhitelist)
	t.Run("Connections", testConnectionsInsert)
	t.Run("Connections", testConnectionsInsertWhitelist)
	t.Run("Dials", testDialsInsert)
	t.Run("Dials", testDialsInsertWhitelist)
	t.Run("FindNodes", testFindNodesInsert)
	t.Run("FindNodes", testFindNodesInsertWhitelist)
	t.Run("IPAddresses", testIPAddressesInsert)
	t.Run("IPAddresses", testIPAddressesInsertWhitelist)
	t.Run("MultiAddresses", testMultiAddressesInsert)
	t.Run("MultiAddresses", testMultiAddressesInsertWhitelist)
	t.Run("PeerLogs", testPeerLogsInsert)
	t.Run("PeerLogs", testPeerLogsInsertWhitelist)
	t.Run("PeerStates", testPeerStatesInsert)
	t.Run("PeerStates", testPeerStatesInsertWhitelist)
	t.Run("Peers", testPeersInsert)
	t.Run("Peers", testPeersInsertWhitelist)
	t.Run("Provides", testProvidesInsert)
	t.Run("Provides", testProvidesInsertWhitelist)
	t.Run("RoutingTableEntries", testRoutingTableEntriesInsert)
	t.Run("RoutingTableEntries", testRoutingTableEntriesInsertWhitelist)
	t.Run("RoutingTableSnapshots", testRoutingTableSnapshotsInsert)
	t.Run("RoutingTableSnapshots", testRoutingTableSnapshotsInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("CloserPeerToFindNodeUsingFindNode", testCloserPeerToOneFindNodeUsingFindNode)
	t.Run("CloserPeerToPeerUsingPeer", testCloserPeerToOnePeerUsingPeer)
	t.Run("CloserPeerToProvideUsingProvide", testCloserPeerToOneProvideUsingProvide)
	t.Run("ConnectionToPeerUsingLocal", testConnectionToOnePeerUsingLocal)
	t.Run("ConnectionToMultiAddressUsingMultiAddress", testConnectionToOneMultiAddressUsingMultiAddress)
	t.Run("ConnectionToProvideUsingProvide", testConnectionToOneProvideUsingProvide)
	t.Run("ConnectionToPeerUsingRemote", testConnectionToOnePeerUsingRemote)
	t.Run("DialToPeerUsingLocal", testDialToOnePeerUsingLocal)
	t.Run("DialToMultiAddressUsingMultiAddress", testDialToOneMultiAddressUsingMultiAddress)
	t.Run("DialToProvideUsingProvide", testDialToOneProvideUsingProvide)
	t.Run("DialToPeerUsingRemote", testDialToOnePeerUsingRemote)
	t.Run("FindNodeToPeerUsingLocal", testFindNodeToOnePeerUsingLocal)
	t.Run("FindNodeToProvideUsingProvide", testFindNodeToOneProvideUsingProvide)
	t.Run("FindNodeToPeerUsingRemote", testFindNodeToOnePeerUsingRemote)
	t.Run("PeerLogToPeerUsingPeer", testPeerLogToOnePeerUsingPeer)
	t.Run("PeerStateToPeerUsingPeer", testPeerStateToOnePeerUsingPeer)
	t.Run("PeerStateToProvideUsingProvide", testPeerStateToOneProvideUsingProvide)
	t.Run("PeerStateToPeerUsingReferrer", testPeerStateToOnePeerUsingReferrer)
	t.Run("ProvideToPeerUsingProvider", testProvideToOnePeerUsingProvider)
	t.Run("RoutingTableEntryToPeerUsingPeer", testRoutingTableEntryToOnePeerUsingPeer)
	t.Run("RoutingTableEntryToRoutingTableSnapshotUsingRoutingTableSnapshot", testRoutingTableEntryToOneRoutingTableSnapshotUsingRoutingTableSnapshot)
	t.Run("RoutingTableSnapshotToPeerUsingPeer", testRoutingTableSnapshotToOnePeerUsingPeer)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("FindNodeToCloserPeers", testFindNodeToManyCloserPeers)
	t.Run("IPAddressToMultiAddresses", testIPAddressToManyMultiAddresses)
	t.Run("MultiAddressToConnections", testMultiAddressToManyConnections)
	t.Run("MultiAddressToDials", testMultiAddressToManyDials)
	t.Run("MultiAddressToIPAddresses", testMultiAddressToManyIPAddresses)
	t.Run("PeerToCloserPeers", testPeerToManyCloserPeers)
	t.Run("PeerToLocalConnections", testPeerToManyLocalConnections)
	t.Run("PeerToRemoteConnections", testPeerToManyRemoteConnections)
	t.Run("PeerToLocalDials", testPeerToManyLocalDials)
	t.Run("PeerToRemoteDials", testPeerToManyRemoteDials)
	t.Run("PeerToLocalFindNodes", testPeerToManyLocalFindNodes)
	t.Run("PeerToRemoteFindNodes", testPeerToManyRemoteFindNodes)
	t.Run("PeerToPeerLogs", testPeerToManyPeerLogs)
	t.Run("PeerToPeerStates", testPeerToManyPeerStates)
	t.Run("PeerToReferrerPeerStates", testPeerToManyReferrerPeerStates)
	t.Run("PeerToProviderProvides", testPeerToManyProviderProvides)
	t.Run("PeerToRoutingTableEntries", testPeerToManyRoutingTableEntries)
	t.Run("PeerToRoutingTableSnapshots", testPeerToManyRoutingTableSnapshots)
	t.Run("ProvideToCloserPeers", testProvideToManyCloserPeers)
	t.Run("ProvideToConnections", testProvideToManyConnections)
	t.Run("ProvideToDials", testProvideToManyDials)
	t.Run("ProvideToFindNodes", testProvideToManyFindNodes)
	t.Run("ProvideToPeerStates", testProvideToManyPeerStates)
	t.Run("RoutingTableSnapshotToRoutingTableEntries", testRoutingTableSnapshotToManyRoutingTableEntries)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("CloserPeerToFindNodeUsingCloserPeers", testCloserPeerToOneSetOpFindNodeUsingFindNode)
	t.Run("CloserPeerToPeerUsingCloserPeers", testCloserPeerToOneSetOpPeerUsingPeer)
	t.Run("CloserPeerToProvideUsingCloserPeers", testCloserPeerToOneSetOpProvideUsingProvide)
	t.Run("ConnectionToPeerUsingLocalConnections", testConnectionToOneSetOpPeerUsingLocal)
	t.Run("ConnectionToMultiAddressUsingConnections", testConnectionToOneSetOpMultiAddressUsingMultiAddress)
	t.Run("ConnectionToProvideUsingConnections", testConnectionToOneSetOpProvideUsingProvide)
	t.Run("ConnectionToPeerUsingRemoteConnections", testConnectionToOneSetOpPeerUsingRemote)
	t.Run("DialToPeerUsingLocalDials", testDialToOneSetOpPeerUsingLocal)
	t.Run("DialToMultiAddressUsingDials", testDialToOneSetOpMultiAddressUsingMultiAddress)
	t.Run("DialToProvideUsingDials", testDialToOneSetOpProvideUsingProvide)
	t.Run("DialToPeerUsingRemoteDials", testDialToOneSetOpPeerUsingRemote)
	t.Run("FindNodeToPeerUsingLocalFindNodes", testFindNodeToOneSetOpPeerUsingLocal)
	t.Run("FindNodeToProvideUsingFindNodes", testFindNodeToOneSetOpProvideUsingProvide)
	t.Run("FindNodeToPeerUsingRemoteFindNodes", testFindNodeToOneSetOpPeerUsingRemote)
	t.Run("PeerLogToPeerUsingPeerLogs", testPeerLogToOneSetOpPeerUsingPeer)
	t.Run("PeerStateToPeerUsingPeerStates", testPeerStateToOneSetOpPeerUsingPeer)
	t.Run("PeerStateToProvideUsingPeerStates", testPeerStateToOneSetOpProvideUsingProvide)
	t.Run("PeerStateToPeerUsingReferrerPeerStates", testPeerStateToOneSetOpPeerUsingReferrer)
	t.Run("ProvideToPeerUsingProviderProvides", testProvideToOneSetOpPeerUsingProvider)
	t.Run("RoutingTableEntryToPeerUsingRoutingTableEntries", testRoutingTableEntryToOneSetOpPeerUsingPeer)
	t.Run("RoutingTableEntryToRoutingTableSnapshotUsingRoutingTableEntries", testRoutingTableEntryToOneSetOpRoutingTableSnapshotUsingRoutingTableSnapshot)
	t.Run("RoutingTableSnapshotToPeerUsingRoutingTableSnapshots", testRoutingTableSnapshotToOneSetOpPeerUsingPeer)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("FindNodeToCloserPeers", testFindNodeToManyAddOpCloserPeers)
	t.Run("IPAddressToMultiAddresses", testIPAddressToManyAddOpMultiAddresses)
	t.Run("MultiAddressToConnections", testMultiAddressToManyAddOpConnections)
	t.Run("MultiAddressToDials", testMultiAddressToManyAddOpDials)
	t.Run("MultiAddressToIPAddresses", testMultiAddressToManyAddOpIPAddresses)
	t.Run("PeerToCloserPeers", testPeerToManyAddOpCloserPeers)
	t.Run("PeerToLocalConnections", testPeerToManyAddOpLocalConnections)
	t.Run("PeerToRemoteConnections", testPeerToManyAddOpRemoteConnections)
	t.Run("PeerToLocalDials", testPeerToManyAddOpLocalDials)
	t.Run("PeerToRemoteDials", testPeerToManyAddOpRemoteDials)
	t.Run("PeerToLocalFindNodes", testPeerToManyAddOpLocalFindNodes)
	t.Run("PeerToRemoteFindNodes", testPeerToManyAddOpRemoteFindNodes)
	t.Run("PeerToPeerLogs", testPeerToManyAddOpPeerLogs)
	t.Run("PeerToPeerStates", testPeerToManyAddOpPeerStates)
	t.Run("PeerToReferrerPeerStates", testPeerToManyAddOpReferrerPeerStates)
	t.Run("PeerToProviderProvides", testPeerToManyAddOpProviderProvides)
	t.Run("PeerToRoutingTableEntries", testPeerToManyAddOpRoutingTableEntries)
	t.Run("PeerToRoutingTableSnapshots", testPeerToManyAddOpRoutingTableSnapshots)
	t.Run("ProvideToCloserPeers", testProvideToManyAddOpCloserPeers)
	t.Run("ProvideToConnections", testProvideToManyAddOpConnections)
	t.Run("ProvideToDials", testProvideToManyAddOpDials)
	t.Run("ProvideToFindNodes", testProvideToManyAddOpFindNodes)
	t.Run("ProvideToPeerStates", testProvideToManyAddOpPeerStates)
	t.Run("RoutingTableSnapshotToRoutingTableEntries", testRoutingTableSnapshotToManyAddOpRoutingTableEntries)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("IPAddressToMultiAddresses", testIPAddressToManySetOpMultiAddresses)
	t.Run("MultiAddressToIPAddresses", testMultiAddressToManySetOpIPAddresses)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("IPAddressToMultiAddresses", testIPAddressToManyRemoveOpMultiAddresses)
	t.Run("MultiAddressToIPAddresses", testMultiAddressToManyRemoveOpIPAddresses)
}

func TestReload(t *testing.T) {
	t.Run("CloserPeers", testCloserPeersReload)
	t.Run("Connections", testConnectionsReload)
	t.Run("Dials", testDialsReload)
	t.Run("FindNodes", testFindNodesReload)
	t.Run("IPAddresses", testIPAddressesReload)
	t.Run("MultiAddresses", testMultiAddressesReload)
	t.Run("PeerLogs", testPeerLogsReload)
	t.Run("PeerStates", testPeerStatesReload)
	t.Run("Peers", testPeersReload)
	t.Run("Provides", testProvidesReload)
	t.Run("RoutingTableEntries", testRoutingTableEntriesReload)
	t.Run("RoutingTableSnapshots", testRoutingTableSnapshotsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("CloserPeers", testCloserPeersReloadAll)
	t.Run("Connections", testConnectionsReloadAll)
	t.Run("Dials", testDialsReloadAll)
	t.Run("FindNodes", testFindNodesReloadAll)
	t.Run("IPAddresses", testIPAddressesReloadAll)
	t.Run("MultiAddresses", testMultiAddressesReloadAll)
	t.Run("PeerLogs", testPeerLogsReloadAll)
	t.Run("PeerStates", testPeerStatesReloadAll)
	t.Run("Peers", testPeersReloadAll)
	t.Run("Provides", testProvidesReloadAll)
	t.Run("RoutingTableEntries", testRoutingTableEntriesReloadAll)
	t.Run("RoutingTableSnapshots", testRoutingTableSnapshotsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("CloserPeers", testCloserPeersSelect)
	t.Run("Connections", testConnectionsSelect)
	t.Run("Dials", testDialsSelect)
	t.Run("FindNodes", testFindNodesSelect)
	t.Run("IPAddresses", testIPAddressesSelect)
	t.Run("MultiAddresses", testMultiAddressesSelect)
	t.Run("PeerLogs", testPeerLogsSelect)
	t.Run("PeerStates", testPeerStatesSelect)
	t.Run("Peers", testPeersSelect)
	t.Run("Provides", testProvidesSelect)
	t.Run("RoutingTableEntries", testRoutingTableEntriesSelect)
	t.Run("RoutingTableSnapshots", testRoutingTableSnapshotsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("CloserPeers", testCloserPeersUpdate)
	t.Run("Connections", testConnectionsUpdate)
	t.Run("Dials", testDialsUpdate)
	t.Run("FindNodes", testFindNodesUpdate)
	t.Run("IPAddresses", testIPAddressesUpdate)
	t.Run("MultiAddresses", testMultiAddressesUpdate)
	t.Run("PeerLogs", testPeerLogsUpdate)
	t.Run("PeerStates", testPeerStatesUpdate)
	t.Run("Peers", testPeersUpdate)
	t.Run("Provides", testProvidesUpdate)
	t.Run("RoutingTableEntries", testRoutingTableEntriesUpdate)
	t.Run("RoutingTableSnapshots", testRoutingTableSnapshotsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("CloserPeers", testCloserPeersSliceUpdateAll)
	t.Run("Connections", testConnectionsSliceUpdateAll)
	t.Run("Dials", testDialsSliceUpdateAll)
	t.Run("FindNodes", testFindNodesSliceUpdateAll)
	t.Run("IPAddresses", testIPAddressesSliceUpdateAll)
	t.Run("MultiAddresses", testMultiAddressesSliceUpdateAll)
	t.Run("PeerLogs", testPeerLogsSliceUpdateAll)
	t.Run("PeerStates", testPeerStatesSliceUpdateAll)
	t.Run("Peers", testPeersSliceUpdateAll)
	t.Run("Provides", testProvidesSliceUpdateAll)
	t.Run("RoutingTableEntries", testRoutingTableEntriesSliceUpdateAll)
	t.Run("RoutingTableSnapshots", testRoutingTableSnapshotsSliceUpdateAll)
}
