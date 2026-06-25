#include <iostream>
using namespace std;

int main()
{
    int arr[] = {4, 2, 7, 8, 1, 2, 5};

    int target = 8;
    int targetIdx = -1;

    for (int i = 0; i < 7; i++)
    {
        if (arr[i] == target)
        {
            targetIdx = i;
            break;
        }
    }
    cout << targetIdx << endl;

    return 0;
}