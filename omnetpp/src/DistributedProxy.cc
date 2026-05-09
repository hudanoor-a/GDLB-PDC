#include <omnetpp.h>
#include <vector>
#include <cmath>
using namespace omnetpp;

class DistributedProxy : public cSimpleModule {
  private:
    std::vector<double> auth;
    std::vector<int> localSent;
    std::vector<simtime_t> lastUpdate;
    double penaltyW;
    double decayAlpha;
  protected:
    void initialize() override {
        int n = gateSize("out");
        auth.assign(n, 0.0);
        localSent.assign(n, 0);
        lastUpdate.assign(n, simTime());
        penaltyW = par("penaltyW").doubleValue();
        decayAlpha = par("decayAlpha").doubleValue();
    }
    int pickBackend() {
        int best = 0;
        double bestScore = 1e100;
        for (int i=0; i<gateSize("out"); i++) {
            double age = (simTime() - lastUpdate[i]).dbl();
            double score = auth[i] + penaltyW * localSent[i] * std::exp(-decayAlpha * age);
            if (score < bestScore) { bestScore = score; best = i; }
        }
        return best;
    }
    void handleMessage(cMessage *msg) override {
        int b = pickBackend();
        localSent[b]++;
        send(msg, "out", b);
    }
};
Define_Module(DistributedProxy);
