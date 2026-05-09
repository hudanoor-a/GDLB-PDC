package gdlb;

simple ClientGenerator
{
    parameters:
        double sendInterval @unit(s);
    gates:
        output out;
}

simple DistributedProxy
{
    parameters:
        string algorithm;
        double penaltyW;
        double decayAlpha;
    gates:
        input in;
        output out[];
}

simple BackendNode
{
    parameters:
        double serviceRate;
    gates:
        input in;
        output telemetryOut;
}

simple MetricsCollector
{
    gates:
        input in[];
}

network GDLBNetwork
{
    parameters:
        int numProxies = default(4);
        int numBackends = default(8);
    submodules:
        client: ClientGenerator;
        proxy[numProxies]: DistributedProxy;
        backend[numBackends]: BackendNode;
        metrics: MetricsCollector;
    connections allowunconnected:
        for i=0..numProxies-1 {
            client.out --> proxy[i].in;
        }
        for i=0..numProxies-1, j=0..numBackends-1 {
            proxy[i].out++ --> { delay = uniform(1ms,5ms); } --> backend[j].in;
        }
}
