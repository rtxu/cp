class ThroneInheritance {
    struct Vertex {
        string name;
        bool dead = false;
        vector<const Vertex*> children;
        
        explicit Vertex(const string& _name): name(_name) {}
        
        void addChild(const Vertex& child) {
            children.push_back(&child);
        }
    };
    
    unordered_map<string, Vertex> g;
    
    void dfs(const Vertex& root, vector<string>& result) {
        if (!root.dead) {
            result.push_back(root.name);
        }
        for (auto child: root.children) {
            dfs(*child, result);
        }
    }
public:
    ThroneInheritance(string kingName) {
        g.emplace(kingName, kingName);
    }
    
    void birth(string parentName, string childName) {
        g.emplace(childName, childName);
        g.at(parentName).addChild(g.at(childName));
    }
    
    void death(string name) {
        g.at(name).dead = true;
    }
    
    vector<string> getInheritanceOrder() {
        vector<string> result;
        dfs(g.at("king"), result);
        return result;
    }
};

/**
 * Your ThroneInheritance object will be instantiated and called as such:
 * ThroneInheritance* obj = new ThroneInheritance(kingName);
 * obj->birth(parentName,childName);
 * obj->death(name);
 * vector<string> param_3 = obj->getInheritanceOrder();
 */
