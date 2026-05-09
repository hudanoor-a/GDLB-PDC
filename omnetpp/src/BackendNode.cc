#include <omnetpp.h>
using namespace omnetpp;

class BackendNode : public cSimpleModule {
  private:
    cQueue queue;
    bool busy = false;
    double serviceRate;
  protected:
    void initialize() override { serviceRate = par("serviceRate").doubleValue(); }
    void handleMessage(cMessage *msg) override {
        queue.insert(msg);
        if (!busy) scheduleAt(simTime(), new cMessage("service"));
    }
};
Define_Module(BackendNode);
