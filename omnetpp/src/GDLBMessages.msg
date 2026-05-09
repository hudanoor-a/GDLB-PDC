packet RequestMsg
{
    int requestId;
    int proxyId;
    simtime_t createdAt;
}

packet TelemetryMsg
{
    int backendId;
    int queueDepth;
    int requestsInFlight;
    simtime_t timestamp;
}
