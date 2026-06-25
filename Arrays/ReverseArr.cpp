#include <iostream>
using namespace std;

int main()
{
    int arr[] = {4, 2, 7, 8, 1, 2, 5};
    int size = 7;
    int head = 0, tail = size - 1;

    for (int i = 0; i < size; i++)
    {
        if (head >= tail)
        {
            break;
        }
        swap(arr[head], arr[tail]);
        cout << "Head " << arr[head] << endl;
        cout << "Tail " << arr[tail] << endl;
        head++;
        tail--;
    }

    cout << "After swap";

    for (int i = 0; i < size; i++)
    {
        cout << arr[i] << endl;
    }

    return 0;
}