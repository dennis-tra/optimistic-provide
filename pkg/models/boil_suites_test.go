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
	t.Run("AddProviderRPCS", testAddProviderRPCS)
	t.Run("CloserPeers", testCloserPeers)
	t.Run("Connections", testConnections)
	t.Run("Dials", testDials)
	t.Run("FindNodesRPCS", testFindNodesRPCS)
	t.Run("GetProvidersRPCS", testGetProvidersRPCS)
	t.Run("Hosts", testHosts)
	t.Run("IPAddresses", testIPAddresses)
	t.Run("MultiAddresses", testMultiAddresses)
	t.Run("PeerLogs", testPeerLogs)
	t.Run("PeerStates", testPeerStates)
	t.Run("Peers", testPeers)
	t.Run("ProviderPeers", testProviderPeers)
	t.Run("Provides", testProvides)
	t.Run("Retrievals", testRetrievals)
	t.Run("RoutingTableEntries", testRoutingTableEntries)
	t.Run("RoutingTableSnapshots", testRoutingTableSnapshots)
}

func TestDelete(t *testing.T) {
	t.Run("AddProviderRPCS", testAddProviderRPCSDelete)
	t.Run("CloserPeers", testCloserPeersDelete)
	t.Run("Connections", testConnectionsDelete)
	t.Run("Dials", testDialsDelete)
	t.Run("FindNodesRPCS", testFindNodesRPCSDelete)
	t.Run("GetProvidersRPCS", testGetProvidersRPCSDelete)
	t.Run("Hosts", testHostsDelete)
	t.Run("IPAddresses", testIPAddressesDelete)
	t.Run("MultiAddresses", testMultiAddressesDelete)
	t.Run("PeerLogs", testPeerLogsDelete)
	t.Run("PeerStates", testPeerStatesDelete)
	t.Run("Peers", testPeersDelete)
	t.Run("ProviderPeers", testProviderPeersDelete)
	t.Run("Provides", testProvidesDelete)
	t.Run("Retrievals", testRetrievalsDelete)
	t.Run("RoutingTableEntries", testRoutingTableEntriesDelete)
	t.Run("RoutingTableSnapshots", testRoutingTableSnapshotsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("AddProviderRPCS", testAddProviderRPCSQueryDeleteAll)
	t.Run("CloserPeers", testCloserPeersQueryDeleteAll)
	t.Run("Connections", testConnectionsQueryDeleteAll)
	t.Run("Dials", testDialsQueryDeleteAll)
	t.Run("FindNodesRPCS", testFindNodesRPCSQueryDeleteAll)
	t.Run("GetProvidersRPCS", testGetProvidersRPCSQueryDeleteAll)
	t.Run("Hosts", testHostsQueryDeleteAll)
	t.Run("IPAddresses", testIPAddressesQueryDeleteAll)
	t.Run("MultiAddresses", testMultiAddressesQueryDeleteAll)
	t.Run("PeerLogs", testPeerLogsQueryDeleteAll)
	t.Run("PeerStates", testPeerStatesQueryDeleteAll)
	t.Run("Peers", testPeersQueryDeleteAll)
	t.Run("ProviderPeers", testProviderPeersQueryDeleteAll)
	t.Run("Provides", testProvidesQueryDeleteAll)
	t.Run("Retrievals", testRetrievalsQueryDeleteAll)
	t.Run("RoutingTableEntries", testRoutingTableEntriesQueryDeleteAll)
	t.Run("RoutingTableSnapshots", testRoutingTableSnapshotsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("AddProviderRPCS", testAddProviderRPCSSliceDeleteAll)
	t.Run("CloserPeers", testCloserPeersSliceDeleteAll)
	t.Run("Connections", testConnectionsSliceDeleteAll)
	t.Run("Dials", testDialsSliceDeleteAll)
	t.Run("FindNodesRPCS", testFindNodesRPCSSliceDeleteAll)
	t.Run("GetProvidersRPCS", testGetProvidersRPCSSliceDeleteAll)
	t.Run("Hosts", testHostsSliceDeleteAll)
	t.Run("IPAddresses", testIPAddressesSliceDeleteAll)
	t.Run("MultiAddresses", testMultiAddressesSliceDeleteAll)
	t.Run("PeerLogs", testPeerLogsSliceDeleteAll)
	t.Run("PeerStates", testPeerStatesSliceDeleteAll)
	t.Run("Peers", testPeersSliceDeleteAll)
	t.Run("ProviderPeers", testProviderPeersSliceDeleteAll)
	t.Run("Provides", testProvidesSliceDeleteAll)
	t.Run("Retrievals", testRetrievalsSliceDeleteAll)
	t.Run("RoutingTableEntries", testRoutingTableEntriesSliceDeleteAll)
	t.Run("RoutingTableSnapshots", testRoutingTableSnapshotsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("AddProviderRPCS", testAddProviderRPCSExists)
	t.Run("CloserPeers", testCloserPeersExists)
	t.Run("Connections", testConnectionsExists)
	t.Run("Dials", testDialsExists)
	t.Run("FindNodesRPCS", testFindNodesRPCSExists)
	t.Run("GetProvidersRPCS", testGetProvidersRPCSExists)
	t.Run("Hosts", testHostsExists)
	t.Run("IPAddresses", testIPAddressesExists)
	t.Run("MultiAddresses", testMultiAddressesExists)
	t.Run("PeerLogs", testPeerLogsExists)
	t.Run("PeerStates", testPeerStatesExists)
	t.Run("Peers", testPeersExists)
	t.Run("ProviderPeers", testProviderPeersExists)
	t.Run("Provides", testProvidesExists)
	t.Run("Retrievals", testRetrievalsExists)
	t.Run("RoutingTableEntries", testRoutingTableEntriesExists)
	t.Run("RoutingTableSnapshots", testRoutingTableSnapshotsExists)
}

func TestFind(t *testing.T) {
	t.Run("AddProviderRPCS", testAddProviderRPCSFind)
	t.Run("CloserPeers", testCloserPeersFind)
	t.Run("Connections", testConnectionsFind)
	t.Run("Dials", testDialsFind)
	t.Run("FindNodesRPCS", testFindNodesRPCSFind)
	t.Run("GetProvidersRPCS", testGetProvidersRPCSFind)
	t.Run("Hosts", testHostsFind)
	t.Run("IPAddresses", testIPAddressesFind)
	t.Run("MultiAddresses", testMultiAddressesFind)
	t.Run("PeerLogs", testPeerLogsFind)
	t.Run("PeerStates", testPeerStatesFind)
	t.Run("Peers", testPeersFind)
	t.Run("ProviderPeers", testProviderPeersFind)
	t.Run("Provides", testProvidesFind)
	t.Run("Retrievals", testRetrievalsFind)
	t.Run("RoutingTableEntries", testRoutingTableEntriesFind)
	t.Run("RoutingTableSnapshots", testRoutingTableSnapshotsFind)
}

func TestBind(t *testing.T) {
	t.Run("AddProviderRPCS", testAddProviderRPCSBind)
	t.Run("CloserPeers", testCloserPeersBind)
	t.Run("Connections", testConnectionsBind)
	t.Run("Dials", testDialsBind)
	t.Run("FindNodesRPCS", testFindNodesRPCSBind)
	t.Run("GetProvidersRPCS", testGetProvidersRPCSBind)
	t.Run("Hosts", testHostsBind)
	t.Run("IPAddresses", testIPAddressesBind)
	t.Run("MultiAddresses", testMultiAddressesBind)
	t.Run("PeerLogs", testPeerLogsBind)
	t.Run("PeerStates", testPeerStatesBind)
	t.Run("Peers", testPeersBind)
	t.Run("ProviderPeers", testProviderPeersBind)
	t.Run("Provides", testProvidesBind)
	t.Run("Retrievals", testRetrievalsBind)
	t.Run("RoutingTableEntries", testRoutingTableEntriesBind)
	t.Run("RoutingTableSnapshots", testRoutingTableSnapshotsBind)
}

func TestOne(t *testing.T) {
	t.Run("AddProviderRPCS", testAddProviderRPCSOne)
	t.Run("CloserPeers", testCloserPeersOne)
	t.Run("Connections", testConnectionsOne)
	t.Run("Dials", testDialsOne)
	t.Run("FindNodesRPCS", testFindNodesRPCSOne)
	t.Run("GetProvidersRPCS", testGetProvidersRPCSOne)
	t.Run("Hosts", testHostsOne)
	t.Run("IPAddresses", testIPAddressesOne)
	t.Run("MultiAddresses", testMultiAddressesOne)
	t.Run("PeerLogs", testPeerLogsOne)
	t.Run("PeerStates", testPeerStatesOne)
	t.Run("Peers", testPeersOne)
	t.Run("ProviderPeers", testProviderPeersOne)
	t.Run("Provides", testProvidesOne)
	t.Run("Retrievals", testRetrievalsOne)
	t.Run("RoutingTableEntries", testRoutingTableEntriesOne)
	t.Run("RoutingTableSnapshots", testRoutingTableSnapshotsOne)
}

func TestAll(t *testing.T) {
	t.Run("AddProviderRPCS", testAddProviderRPCSAll)
	t.Run("CloserPeers", testCloserPeersAll)
	t.Run("Connections", testConnectionsAll)
	t.Run("Dials", testDialsAll)
	t.Run("FindNodesRPCS", testFindNodesRPCSAll)
	t.Run("GetProvidersRPCS", testGetProvidersRPCSAll)
	t.Run("Hosts", testHostsAll)
	t.Run("IPAddresses", testIPAddressesAll)
	t.Run("MultiAddresses", testMultiAddressesAll)
	t.Run("PeerLogs", testPeerLogsAll)
	t.Run("PeerStates", testPeerStatesAll)
	t.Run("Peers", testPeersAll)
	t.Run("ProviderPeers", testProviderPeersAll)
	t.Run("Provides", testProvidesAll)
	t.Run("Retrievals", testRetrievalsAll)
	t.Run("RoutingTableEntries", testRoutingTableEntriesAll)
	t.Run("RoutingTableSnapshots", testRoutingTableSnapshotsAll)
}

func TestCount(t *testing.T) {
	t.Run("AddProviderRPCS", testAddProviderRPCSCount)
	t.Run("CloserPeers", testCloserPeersCount)
	t.Run("Connections", testConnectionsCount)
	t.Run("Dials", testDialsCount)
	t.Run("FindNodesRPCS", testFindNodesRPCSCount)
	t.Run("GetProvidersRPCS", testGetProvidersRPCSCount)
	t.Run("Hosts", testHostsCount)
	t.Run("IPAddresses", testIPAddressesCount)
	t.Run("MultiAddresses", testMultiAddressesCount)
	t.Run("PeerLogs", testPeerLogsCount)
	t.Run("PeerStates", testPeerStatesCount)
	t.Run("Peers", testPeersCount)
	t.Run("ProviderPeers", testProviderPeersCount)
	t.Run("Provides", testProvidesCount)
	t.Run("Retrievals", testRetrievalsCount)
	t.Run("RoutingTableEntries", testRoutingTableEntriesCount)
	t.Run("RoutingTableSnapshots", testRoutingTableSnapshotsCount)
}

func TestHooks(t *testing.T) {
	t.Run("AddProviderRPCS", testAddProviderRPCSHooks)
	t.Run("CloserPeers", testCloserPeersHooks)
	t.Run("Connections", testConnectionsHooks)
	t.Run("Dials", testDialsHooks)
	t.Run("FindNodesRPCS", testFindNodesRPCSHooks)
	t.Run("GetProvidersRPCS", testGetProvidersRPCSHooks)
	t.Run("Hosts", testHostsHooks)
	t.Run("IPAddresses", testIPAddressesHooks)
	t.Run("MultiAddresses", testMultiAddressesHooks)
	t.Run("PeerLogs", testPeerLogsHooks)
	t.Run("PeerStates", testPeerStatesHooks)
	t.Run("Peers", testPeersHooks)
	t.Run("ProviderPeers", testProviderPeersHooks)
	t.Run("Provides", testProvidesHooks)
	t.Run("Retrievals", testRetrievalsHooks)
	t.Run("RoutingTableEntries", testRoutingTableEntriesHooks)
	t.Run("RoutingTableSnapshots", testRoutingTableSnapshotsHooks)
}

func TestInsert(t *testing.T) {
	t.Run("AddProviderRPCS", testAddProviderRPCSInsert)
	t.Run("AddProviderRPCS", testAddProviderRPCSInsertWhitelist)
	t.Run("CloserPeers", testCloserPeersInsert)
	t.Run("CloserPeers", testCloserPeersInsertWhitelist)
	t.Run("Connections", testConnectionsInsert)
	t.Run("Connections", testConnectionsInsertWhitelist)
	t.Run("Dials", testDialsInsert)
	t.Run("Dials", testDialsInsertWhitelist)
	t.Run("FindNodesRPCS", testFindNodesRPCSInsert)
	t.Run("FindNodesRPCS", testFindNodesRPCSInsertWhitelist)
	t.Run("GetProvidersRPCS", testGetProvidersRPCSInsert)
	t.Run("GetProvidersRPCS", testGetProvidersRPCSInsertWhitelist)
	t.Run("Hosts", testHostsInsert)
	t.Run("Hosts", testHostsInsertWhitelist)
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
	t.Run("ProviderPeers", testProviderPeersInsert)
	t.Run("ProviderPeers", testProviderPeersInsertWhitelist)
	t.Run("Provides", testProvidesInsert)
	t.Run("Provides", testProvidesInsertWhitelist)
	t.Run("Retrievals", testRetrievalsInsert)
	t.Run("Retrievals", testRetrievalsInsertWhitelist)
	t.Run("RoutingTableEntries", testRoutingTableEntriesInsert)
	t.Run("RoutingTableEntries", testRoutingTableEntriesInsertWhitelist)
	t.Run("RoutingTableSnapshots", testRoutingTableSnapshotsInsert)
	t.Run("RoutingTableSnapshots", testRoutingTableSnapshotsInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("AddProviderRPCToPeerUsingLocal", testAddProviderRPCToOnePeerUsingLocal)
	t.Run("AddProviderRPCToPeerUsingRemote", testAddProviderRPCToOnePeerUsingRemote)
	t.Run("CloserPeerToFindNodesRPCUsingFindNodeRPC", testCloserPeerToOneFindNodesRPCUsingFindNodeRPC)
	t.Run("CloserPeerToPeerUsingPeer", testCloserPeerToOnePeerUsingPeer)
	t.Run("ConnectionToPeerUsingLocal", testConnectionToOnePeerUsingLocal)
	t.Run("ConnectionToMultiAddressUsingMultiAddress", testConnectionToOneMultiAddressUsingMultiAddress)
	t.Run("ConnectionToPeerUsingRemote", testConnectionToOnePeerUsingRemote)
	t.Run("DialToPeerUsingLocal", testDialToOnePeerUsingLocal)
	t.Run("DialToMultiAddressUsingMultiAddress", testDialToOneMultiAddressUsingMultiAddress)
	t.Run("DialToPeerUsingRemote", testDialToOnePeerUsingRemote)
	t.Run("FindNodesRPCToPeerUsingLocal", testFindNodesRPCToOnePeerUsingLocal)
	t.Run("FindNodesRPCToPeerUsingRemote", testFindNodesRPCToOnePeerUsingRemote)
	t.Run("GetProvidersRPCToPeerUsingLocal", testGetProvidersRPCToOnePeerUsingLocal)
	t.Run("GetProvidersRPCToPeerUsingRemote", testGetProvidersRPCToOnePeerUsingRemote)
	t.Run("HostToPeerUsingPeer", testHostToOnePeerUsingPeer)
	t.Run("PeerLogToPeerUsingPeer", testPeerLogToOnePeerUsingPeer)
	t.Run("PeerStateToPeerUsingPeer", testPeerStateToOnePeerUsingPeer)
	t.Run("PeerStateToPeerUsingReferrer", testPeerStateToOnePeerUsingReferrer)
	t.Run("ProviderPeerToGetProvidersRPCUsingGetProvidersRPC", testProviderPeerToOneGetProvidersRPCUsingGetProvidersRPC)
	t.Run("ProviderPeerToPeerUsingProvider", testProviderPeerToOnePeerUsingProvider)
	t.Run("ProvideToRoutingTableSnapshotUsingFinalRoutingTable", testProvideToOneRoutingTableSnapshotUsingFinalRoutingTable)
	t.Run("ProvideToRoutingTableSnapshotUsingInitialRoutingTable", testProvideToOneRoutingTableSnapshotUsingInitialRoutingTable)
	t.Run("ProvideToPeerUsingProvider", testProvideToOnePeerUsingProvider)
	t.Run("RetrievalToRoutingTableSnapshotUsingFinalRoutingTable", testRetrievalToOneRoutingTableSnapshotUsingFinalRoutingTable)
	t.Run("RetrievalToRoutingTableSnapshotUsingInitialRoutingTable", testRetrievalToOneRoutingTableSnapshotUsingInitialRoutingTable)
	t.Run("RetrievalToPeerUsingRetriever", testRetrievalToOnePeerUsingRetriever)
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
	t.Run("AddProviderRPCToProvides", testAddProviderRPCToManyProvides)
	t.Run("ConnectionToProvides", testConnectionToManyProvides)
	t.Run("ConnectionToRetrievals", testConnectionToManyRetrievals)
	t.Run("DialToProvides", testDialToManyProvides)
	t.Run("DialToRetrievals", testDialToManyRetrievals)
	t.Run("FindNodesRPCToFindNodeRPCCloserPeers", testFindNodesRPCToManyFindNodeRPCCloserPeers)
	t.Run("FindNodesRPCToProvides", testFindNodesRPCToManyProvides)
	t.Run("GetProvidersRPCToProviderPeers", testGetProvidersRPCToManyProviderPeers)
	t.Run("GetProvidersRPCToRetrievals", testGetProvidersRPCToManyRetrievals)
	t.Run("IPAddressToMultiAddresses", testIPAddressToManyMultiAddresses)
	t.Run("MultiAddressToConnections", testMultiAddressToManyConnections)
	t.Run("MultiAddressToDials", testMultiAddressToManyDials)
	t.Run("MultiAddressToIPAddresses", testMultiAddressToManyIPAddresses)
	t.Run("PeerStateToProvides", testPeerStateToManyProvides)
	t.Run("PeerStateToRetrievals", testPeerStateToManyRetrievals)
	t.Run("PeerToLocalAddProviderRPCS", testPeerToManyLocalAddProviderRPCS)
	t.Run("PeerToRemoteAddProviderRPCS", testPeerToManyRemoteAddProviderRPCS)
	t.Run("PeerToCloserPeers", testPeerToManyCloserPeers)
	t.Run("PeerToLocalConnections", testPeerToManyLocalConnections)
	t.Run("PeerToRemoteConnections", testPeerToManyRemoteConnections)
	t.Run("PeerToLocalDials", testPeerToManyLocalDials)
	t.Run("PeerToRemoteDials", testPeerToManyRemoteDials)
	t.Run("PeerToLocalFindNodesRPCS", testPeerToManyLocalFindNodesRPCS)
	t.Run("PeerToRemoteFindNodesRPCS", testPeerToManyRemoteFindNodesRPCS)
	t.Run("PeerToLocalGetProvidersRPCS", testPeerToManyLocalGetProvidersRPCS)
	t.Run("PeerToRemoteGetProvidersRPCS", testPeerToManyRemoteGetProvidersRPCS)
	t.Run("PeerToHosts", testPeerToManyHosts)
	t.Run("PeerToPeerLogs", testPeerToManyPeerLogs)
	t.Run("PeerToPeerStates", testPeerToManyPeerStates)
	t.Run("PeerToReferrerPeerStates", testPeerToManyReferrerPeerStates)
	t.Run("PeerToProviderProviderPeers", testPeerToManyProviderProviderPeers)
	t.Run("PeerToProviderProvides", testPeerToManyProviderProvides)
	t.Run("PeerToRetrieverRetrievals", testPeerToManyRetrieverRetrievals)
	t.Run("PeerToRoutingTableEntries", testPeerToManyRoutingTableEntries)
	t.Run("PeerToRoutingTableSnapshots", testPeerToManyRoutingTableSnapshots)
	t.Run("ProvideToAddProviderRPCS", testProvideToManyAddProviderRPCS)
	t.Run("ProvideToConnections", testProvideToManyConnections)
	t.Run("ProvideToDials", testProvideToManyDials)
	t.Run("ProvideToFindNodesRPCS", testProvideToManyFindNodesRPCS)
	t.Run("ProvideToPeerStates", testProvideToManyPeerStates)
	t.Run("RetrievalToConnections", testRetrievalToManyConnections)
	t.Run("RetrievalToDials", testRetrievalToManyDials)
	t.Run("RetrievalToGetProvidersRPCS", testRetrievalToManyGetProvidersRPCS)
	t.Run("RetrievalToPeerStates", testRetrievalToManyPeerStates)
	t.Run("RoutingTableSnapshotToFinalRoutingTableProvides", testRoutingTableSnapshotToManyFinalRoutingTableProvides)
	t.Run("RoutingTableSnapshotToInitialRoutingTableProvides", testRoutingTableSnapshotToManyInitialRoutingTableProvides)
	t.Run("RoutingTableSnapshotToFinalRoutingTableRetrievals", testRoutingTableSnapshotToManyFinalRoutingTableRetrievals)
	t.Run("RoutingTableSnapshotToInitialRoutingTableRetrievals", testRoutingTableSnapshotToManyInitialRoutingTableRetrievals)
	t.Run("RoutingTableSnapshotToRoutingTableEntries", testRoutingTableSnapshotToManyRoutingTableEntries)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("AddProviderRPCToPeerUsingLocalAddProviderRPCS", testAddProviderRPCToOneSetOpPeerUsingLocal)
	t.Run("AddProviderRPCToPeerUsingRemoteAddProviderRPCS", testAddProviderRPCToOneSetOpPeerUsingRemote)
	t.Run("CloserPeerToFindNodesRPCUsingFindNodeRPCCloserPeers", testCloserPeerToOneSetOpFindNodesRPCUsingFindNodeRPC)
	t.Run("CloserPeerToPeerUsingCloserPeers", testCloserPeerToOneSetOpPeerUsingPeer)
	t.Run("ConnectionToPeerUsingLocalConnections", testConnectionToOneSetOpPeerUsingLocal)
	t.Run("ConnectionToMultiAddressUsingConnections", testConnectionToOneSetOpMultiAddressUsingMultiAddress)
	t.Run("ConnectionToPeerUsingRemoteConnections", testConnectionToOneSetOpPeerUsingRemote)
	t.Run("DialToPeerUsingLocalDials", testDialToOneSetOpPeerUsingLocal)
	t.Run("DialToMultiAddressUsingDials", testDialToOneSetOpMultiAddressUsingMultiAddress)
	t.Run("DialToPeerUsingRemoteDials", testDialToOneSetOpPeerUsingRemote)
	t.Run("FindNodesRPCToPeerUsingLocalFindNodesRPCS", testFindNodesRPCToOneSetOpPeerUsingLocal)
	t.Run("FindNodesRPCToPeerUsingRemoteFindNodesRPCS", testFindNodesRPCToOneSetOpPeerUsingRemote)
	t.Run("GetProvidersRPCToPeerUsingLocalGetProvidersRPCS", testGetProvidersRPCToOneSetOpPeerUsingLocal)
	t.Run("GetProvidersRPCToPeerUsingRemoteGetProvidersRPCS", testGetProvidersRPCToOneSetOpPeerUsingRemote)
	t.Run("HostToPeerUsingHosts", testHostToOneSetOpPeerUsingPeer)
	t.Run("PeerLogToPeerUsingPeerLogs", testPeerLogToOneSetOpPeerUsingPeer)
	t.Run("PeerStateToPeerUsingPeerStates", testPeerStateToOneSetOpPeerUsingPeer)
	t.Run("PeerStateToPeerUsingReferrerPeerStates", testPeerStateToOneSetOpPeerUsingReferrer)
	t.Run("ProviderPeerToGetProvidersRPCUsingProviderPeers", testProviderPeerToOneSetOpGetProvidersRPCUsingGetProvidersRPC)
	t.Run("ProviderPeerToPeerUsingProviderProviderPeers", testProviderPeerToOneSetOpPeerUsingProvider)
	t.Run("ProvideToRoutingTableSnapshotUsingFinalRoutingTableProvides", testProvideToOneSetOpRoutingTableSnapshotUsingFinalRoutingTable)
	t.Run("ProvideToRoutingTableSnapshotUsingInitialRoutingTableProvides", testProvideToOneSetOpRoutingTableSnapshotUsingInitialRoutingTable)
	t.Run("ProvideToPeerUsingProviderProvides", testProvideToOneSetOpPeerUsingProvider)
	t.Run("RetrievalToRoutingTableSnapshotUsingFinalRoutingTableRetrievals", testRetrievalToOneSetOpRoutingTableSnapshotUsingFinalRoutingTable)
	t.Run("RetrievalToRoutingTableSnapshotUsingInitialRoutingTableRetrievals", testRetrievalToOneSetOpRoutingTableSnapshotUsingInitialRoutingTable)
	t.Run("RetrievalToPeerUsingRetrieverRetrievals", testRetrievalToOneSetOpPeerUsingRetriever)
	t.Run("RoutingTableEntryToPeerUsingRoutingTableEntries", testRoutingTableEntryToOneSetOpPeerUsingPeer)
	t.Run("RoutingTableEntryToRoutingTableSnapshotUsingRoutingTableEntries", testRoutingTableEntryToOneSetOpRoutingTableSnapshotUsingRoutingTableSnapshot)
	t.Run("RoutingTableSnapshotToPeerUsingRoutingTableSnapshots", testRoutingTableSnapshotToOneSetOpPeerUsingPeer)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {
	t.Run("ProvideToRoutingTableSnapshotUsingFinalRoutingTableProvides", testProvideToOneRemoveOpRoutingTableSnapshotUsingFinalRoutingTable)
	t.Run("RetrievalToRoutingTableSnapshotUsingFinalRoutingTableRetrievals", testRetrievalToOneRemoveOpRoutingTableSnapshotUsingFinalRoutingTable)
}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("AddProviderRPCToProvides", testAddProviderRPCToManyAddOpProvides)
	t.Run("ConnectionToProvides", testConnectionToManyAddOpProvides)
	t.Run("ConnectionToRetrievals", testConnectionToManyAddOpRetrievals)
	t.Run("DialToProvides", testDialToManyAddOpProvides)
	t.Run("DialToRetrievals", testDialToManyAddOpRetrievals)
	t.Run("FindNodesRPCToFindNodeRPCCloserPeers", testFindNodesRPCToManyAddOpFindNodeRPCCloserPeers)
	t.Run("FindNodesRPCToProvides", testFindNodesRPCToManyAddOpProvides)
	t.Run("GetProvidersRPCToProviderPeers", testGetProvidersRPCToManyAddOpProviderPeers)
	t.Run("GetProvidersRPCToRetrievals", testGetProvidersRPCToManyAddOpRetrievals)
	t.Run("IPAddressToMultiAddresses", testIPAddressToManyAddOpMultiAddresses)
	t.Run("MultiAddressToConnections", testMultiAddressToManyAddOpConnections)
	t.Run("MultiAddressToDials", testMultiAddressToManyAddOpDials)
	t.Run("MultiAddressToIPAddresses", testMultiAddressToManyAddOpIPAddresses)
	t.Run("PeerStateToProvides", testPeerStateToManyAddOpProvides)
	t.Run("PeerStateToRetrievals", testPeerStateToManyAddOpRetrievals)
	t.Run("PeerToLocalAddProviderRPCS", testPeerToManyAddOpLocalAddProviderRPCS)
	t.Run("PeerToRemoteAddProviderRPCS", testPeerToManyAddOpRemoteAddProviderRPCS)
	t.Run("PeerToCloserPeers", testPeerToManyAddOpCloserPeers)
	t.Run("PeerToLocalConnections", testPeerToManyAddOpLocalConnections)
	t.Run("PeerToRemoteConnections", testPeerToManyAddOpRemoteConnections)
	t.Run("PeerToLocalDials", testPeerToManyAddOpLocalDials)
	t.Run("PeerToRemoteDials", testPeerToManyAddOpRemoteDials)
	t.Run("PeerToLocalFindNodesRPCS", testPeerToManyAddOpLocalFindNodesRPCS)
	t.Run("PeerToRemoteFindNodesRPCS", testPeerToManyAddOpRemoteFindNodesRPCS)
	t.Run("PeerToLocalGetProvidersRPCS", testPeerToManyAddOpLocalGetProvidersRPCS)
	t.Run("PeerToRemoteGetProvidersRPCS", testPeerToManyAddOpRemoteGetProvidersRPCS)
	t.Run("PeerToHosts", testPeerToManyAddOpHosts)
	t.Run("PeerToPeerLogs", testPeerToManyAddOpPeerLogs)
	t.Run("PeerToPeerStates", testPeerToManyAddOpPeerStates)
	t.Run("PeerToReferrerPeerStates", testPeerToManyAddOpReferrerPeerStates)
	t.Run("PeerToProviderProviderPeers", testPeerToManyAddOpProviderProviderPeers)
	t.Run("PeerToProviderProvides", testPeerToManyAddOpProviderProvides)
	t.Run("PeerToRetrieverRetrievals", testPeerToManyAddOpRetrieverRetrievals)
	t.Run("PeerToRoutingTableEntries", testPeerToManyAddOpRoutingTableEntries)
	t.Run("PeerToRoutingTableSnapshots", testPeerToManyAddOpRoutingTableSnapshots)
	t.Run("ProvideToAddProviderRPCS", testProvideToManyAddOpAddProviderRPCS)
	t.Run("ProvideToConnections", testProvideToManyAddOpConnections)
	t.Run("ProvideToDials", testProvideToManyAddOpDials)
	t.Run("ProvideToFindNodesRPCS", testProvideToManyAddOpFindNodesRPCS)
	t.Run("ProvideToPeerStates", testProvideToManyAddOpPeerStates)
	t.Run("RetrievalToConnections", testRetrievalToManyAddOpConnections)
	t.Run("RetrievalToDials", testRetrievalToManyAddOpDials)
	t.Run("RetrievalToGetProvidersRPCS", testRetrievalToManyAddOpGetProvidersRPCS)
	t.Run("RetrievalToPeerStates", testRetrievalToManyAddOpPeerStates)
	t.Run("RoutingTableSnapshotToFinalRoutingTableProvides", testRoutingTableSnapshotToManyAddOpFinalRoutingTableProvides)
	t.Run("RoutingTableSnapshotToInitialRoutingTableProvides", testRoutingTableSnapshotToManyAddOpInitialRoutingTableProvides)
	t.Run("RoutingTableSnapshotToFinalRoutingTableRetrievals", testRoutingTableSnapshotToManyAddOpFinalRoutingTableRetrievals)
	t.Run("RoutingTableSnapshotToInitialRoutingTableRetrievals", testRoutingTableSnapshotToManyAddOpInitialRoutingTableRetrievals)
	t.Run("RoutingTableSnapshotToRoutingTableEntries", testRoutingTableSnapshotToManyAddOpRoutingTableEntries)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("AddProviderRPCToProvides", testAddProviderRPCToManySetOpProvides)
	t.Run("ConnectionToProvides", testConnectionToManySetOpProvides)
	t.Run("ConnectionToRetrievals", testConnectionToManySetOpRetrievals)
	t.Run("DialToProvides", testDialToManySetOpProvides)
	t.Run("DialToRetrievals", testDialToManySetOpRetrievals)
	t.Run("FindNodesRPCToProvides", testFindNodesRPCToManySetOpProvides)
	t.Run("GetProvidersRPCToRetrievals", testGetProvidersRPCToManySetOpRetrievals)
	t.Run("IPAddressToMultiAddresses", testIPAddressToManySetOpMultiAddresses)
	t.Run("MultiAddressToIPAddresses", testMultiAddressToManySetOpIPAddresses)
	t.Run("PeerStateToProvides", testPeerStateToManySetOpProvides)
	t.Run("PeerStateToRetrievals", testPeerStateToManySetOpRetrievals)
	t.Run("ProvideToAddProviderRPCS", testProvideToManySetOpAddProviderRPCS)
	t.Run("ProvideToConnections", testProvideToManySetOpConnections)
	t.Run("ProvideToDials", testProvideToManySetOpDials)
	t.Run("ProvideToFindNodesRPCS", testProvideToManySetOpFindNodesRPCS)
	t.Run("ProvideToPeerStates", testProvideToManySetOpPeerStates)
	t.Run("RetrievalToConnections", testRetrievalToManySetOpConnections)
	t.Run("RetrievalToDials", testRetrievalToManySetOpDials)
	t.Run("RetrievalToGetProvidersRPCS", testRetrievalToManySetOpGetProvidersRPCS)
	t.Run("RetrievalToPeerStates", testRetrievalToManySetOpPeerStates)
	t.Run("RoutingTableSnapshotToFinalRoutingTableProvides", testRoutingTableSnapshotToManySetOpFinalRoutingTableProvides)
	t.Run("RoutingTableSnapshotToFinalRoutingTableRetrievals", testRoutingTableSnapshotToManySetOpFinalRoutingTableRetrievals)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("AddProviderRPCToProvides", testAddProviderRPCToManyRemoveOpProvides)
	t.Run("ConnectionToProvides", testConnectionToManyRemoveOpProvides)
	t.Run("ConnectionToRetrievals", testConnectionToManyRemoveOpRetrievals)
	t.Run("DialToProvides", testDialToManyRemoveOpProvides)
	t.Run("DialToRetrievals", testDialToManyRemoveOpRetrievals)
	t.Run("FindNodesRPCToProvides", testFindNodesRPCToManyRemoveOpProvides)
	t.Run("GetProvidersRPCToRetrievals", testGetProvidersRPCToManyRemoveOpRetrievals)
	t.Run("IPAddressToMultiAddresses", testIPAddressToManyRemoveOpMultiAddresses)
	t.Run("MultiAddressToIPAddresses", testMultiAddressToManyRemoveOpIPAddresses)
	t.Run("PeerStateToProvides", testPeerStateToManyRemoveOpProvides)
	t.Run("PeerStateToRetrievals", testPeerStateToManyRemoveOpRetrievals)
	t.Run("ProvideToAddProviderRPCS", testProvideToManyRemoveOpAddProviderRPCS)
	t.Run("ProvideToConnections", testProvideToManyRemoveOpConnections)
	t.Run("ProvideToDials", testProvideToManyRemoveOpDials)
	t.Run("ProvideToFindNodesRPCS", testProvideToManyRemoveOpFindNodesRPCS)
	t.Run("ProvideToPeerStates", testProvideToManyRemoveOpPeerStates)
	t.Run("RetrievalToConnections", testRetrievalToManyRemoveOpConnections)
	t.Run("RetrievalToDials", testRetrievalToManyRemoveOpDials)
	t.Run("RetrievalToGetProvidersRPCS", testRetrievalToManyRemoveOpGetProvidersRPCS)
	t.Run("RetrievalToPeerStates", testRetrievalToManyRemoveOpPeerStates)
	t.Run("RoutingTableSnapshotToFinalRoutingTableProvides", testRoutingTableSnapshotToManyRemoveOpFinalRoutingTableProvides)
	t.Run("RoutingTableSnapshotToFinalRoutingTableRetrievals", testRoutingTableSnapshotToManyRemoveOpFinalRoutingTableRetrievals)
}

func TestReload(t *testing.T) {
	t.Run("AddProviderRPCS", testAddProviderRPCSReload)
	t.Run("CloserPeers", testCloserPeersReload)
	t.Run("Connections", testConnectionsReload)
	t.Run("Dials", testDialsReload)
	t.Run("FindNodesRPCS", testFindNodesRPCSReload)
	t.Run("GetProvidersRPCS", testGetProvidersRPCSReload)
	t.Run("Hosts", testHostsReload)
	t.Run("IPAddresses", testIPAddressesReload)
	t.Run("MultiAddresses", testMultiAddressesReload)
	t.Run("PeerLogs", testPeerLogsReload)
	t.Run("PeerStates", testPeerStatesReload)
	t.Run("Peers", testPeersReload)
	t.Run("ProviderPeers", testProviderPeersReload)
	t.Run("Provides", testProvidesReload)
	t.Run("Retrievals", testRetrievalsReload)
	t.Run("RoutingTableEntries", testRoutingTableEntriesReload)
	t.Run("RoutingTableSnapshots", testRoutingTableSnapshotsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("AddProviderRPCS", testAddProviderRPCSReloadAll)
	t.Run("CloserPeers", testCloserPeersReloadAll)
	t.Run("Connections", testConnectionsReloadAll)
	t.Run("Dials", testDialsReloadAll)
	t.Run("FindNodesRPCS", testFindNodesRPCSReloadAll)
	t.Run("GetProvidersRPCS", testGetProvidersRPCSReloadAll)
	t.Run("Hosts", testHostsReloadAll)
	t.Run("IPAddresses", testIPAddressesReloadAll)
	t.Run("MultiAddresses", testMultiAddressesReloadAll)
	t.Run("PeerLogs", testPeerLogsReloadAll)
	t.Run("PeerStates", testPeerStatesReloadAll)
	t.Run("Peers", testPeersReloadAll)
	t.Run("ProviderPeers", testProviderPeersReloadAll)
	t.Run("Provides", testProvidesReloadAll)
	t.Run("Retrievals", testRetrievalsReloadAll)
	t.Run("RoutingTableEntries", testRoutingTableEntriesReloadAll)
	t.Run("RoutingTableSnapshots", testRoutingTableSnapshotsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("AddProviderRPCS", testAddProviderRPCSSelect)
	t.Run("CloserPeers", testCloserPeersSelect)
	t.Run("Connections", testConnectionsSelect)
	t.Run("Dials", testDialsSelect)
	t.Run("FindNodesRPCS", testFindNodesRPCSSelect)
	t.Run("GetProvidersRPCS", testGetProvidersRPCSSelect)
	t.Run("Hosts", testHostsSelect)
	t.Run("IPAddresses", testIPAddressesSelect)
	t.Run("MultiAddresses", testMultiAddressesSelect)
	t.Run("PeerLogs", testPeerLogsSelect)
	t.Run("PeerStates", testPeerStatesSelect)
	t.Run("Peers", testPeersSelect)
	t.Run("ProviderPeers", testProviderPeersSelect)
	t.Run("Provides", testProvidesSelect)
	t.Run("Retrievals", testRetrievalsSelect)
	t.Run("RoutingTableEntries", testRoutingTableEntriesSelect)
	t.Run("RoutingTableSnapshots", testRoutingTableSnapshotsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("AddProviderRPCS", testAddProviderRPCSUpdate)
	t.Run("CloserPeers", testCloserPeersUpdate)
	t.Run("Connections", testConnectionsUpdate)
	t.Run("Dials", testDialsUpdate)
	t.Run("FindNodesRPCS", testFindNodesRPCSUpdate)
	t.Run("GetProvidersRPCS", testGetProvidersRPCSUpdate)
	t.Run("Hosts", testHostsUpdate)
	t.Run("IPAddresses", testIPAddressesUpdate)
	t.Run("MultiAddresses", testMultiAddressesUpdate)
	t.Run("PeerLogs", testPeerLogsUpdate)
	t.Run("PeerStates", testPeerStatesUpdate)
	t.Run("Peers", testPeersUpdate)
	t.Run("ProviderPeers", testProviderPeersUpdate)
	t.Run("Provides", testProvidesUpdate)
	t.Run("Retrievals", testRetrievalsUpdate)
	t.Run("RoutingTableEntries", testRoutingTableEntriesUpdate)
	t.Run("RoutingTableSnapshots", testRoutingTableSnapshotsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("AddProviderRPCS", testAddProviderRPCSSliceUpdateAll)
	t.Run("CloserPeers", testCloserPeersSliceUpdateAll)
	t.Run("Connections", testConnectionsSliceUpdateAll)
	t.Run("Dials", testDialsSliceUpdateAll)
	t.Run("FindNodesRPCS", testFindNodesRPCSSliceUpdateAll)
	t.Run("GetProvidersRPCS", testGetProvidersRPCSSliceUpdateAll)
	t.Run("Hosts", testHostsSliceUpdateAll)
	t.Run("IPAddresses", testIPAddressesSliceUpdateAll)
	t.Run("MultiAddresses", testMultiAddressesSliceUpdateAll)
	t.Run("PeerLogs", testPeerLogsSliceUpdateAll)
	t.Run("PeerStates", testPeerStatesSliceUpdateAll)
	t.Run("Peers", testPeersSliceUpdateAll)
	t.Run("ProviderPeers", testProviderPeersSliceUpdateAll)
	t.Run("Provides", testProvidesSliceUpdateAll)
	t.Run("Retrievals", testRetrievalsSliceUpdateAll)
	t.Run("RoutingTableEntries", testRoutingTableEntriesSliceUpdateAll)
	t.Run("RoutingTableSnapshots", testRoutingTableSnapshotsSliceUpdateAll)
}
