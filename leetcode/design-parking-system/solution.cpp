class ParkingSystem {
    array<int, 3> mAvailable;
public:
    ParkingSystem(int big, int medium, int small) {
        mAvailable[0] = big;
        mAvailable[1] = medium;
        mAvailable[2] = small;
    }
    
    bool addCar(int carType) {
        carType--;
        if (mAvailable[carType] > 0) {
            mAvailable[carType]--;
            return true;
        } else {
            return false;
        }
    }
};

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * ParkingSystem* obj = new ParkingSystem(big, medium, small);
 * bool param_1 = obj->addCar(carType);
 */
