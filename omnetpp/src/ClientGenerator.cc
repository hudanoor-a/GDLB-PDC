#include <omnetpp.h>
using namespace omnetpp;

class ClientGenerator : public cSimpleModule {
  private:
    int seq = 0;
  protected:
    void initialize() override { scheduleAt(simTime(), new cMessage("tick")); }
    void handleMessage(cMessage *msg) override {
        delete msg;
        auto *req = new cMessage((std::string("req-")+std::to_string(seq++)).c_str());
        send(req, "out");
        scheduleAt(simTime() + par("sendInterval"), new cMessage("tick"));
    }
};
Define_Module(ClientGenerator);
