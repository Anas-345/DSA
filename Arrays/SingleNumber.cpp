#include <iostream>
#include <vector>
using namespace std;

int main()
{
    vector<int> nums = {4, 1, 2, 1, 2};
    int size = nums.size();

    for (int i = 0; i < size; i++)
    {
        bool check = false;
        for (int j = 0; j < size; j++)
        {
            if (i != j && nums[i] == nums[j])
            {
                check = true;
                break;
            }
        }
        if (!check)
        {
            cout << nums[i] << endl;
        }
    }

    return 0;
}