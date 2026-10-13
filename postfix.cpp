#include <iostream>
using namespace std;

char opStack[50] = {'('};
int top = 0, ansStack[50], ansTop = -1;

void ansPush(int num)
{
    ansTop++;
    ansStack[ansTop] = num;
}

int ansPop()
{
    return ansStack[ansTop--];
}

void opPush(char x)
{
    top++;
    opStack[top] = x;
}

char opPop()
{
    return opStack[top--];
}

bool isNumber(char num)
{
    if (num >= 48 && num <= 57)
    {
        return true;
    }
    return false;
}

bool isParanthesis(char paran)
{
    if (paran == 40 || paran == 41)
    {
        return true;
    }
    return false;
}

bool isOperator(char op)
{
    if (op == '+' || op == '-' || op == '*' || op == '/')
    {
        return true;
    }
    return false;
}

int getPriority(char op)
{
    if (op == '+' || op == '-')
        return 1;
    return 2;
}

int operations(int x, int y, char op)
{
    switch (op)
    {
    case '+':
        return x + y;
    case '-':
        return x - y;
    case '*':
        return x * y;
    default:
        return x / y;
    }
}

main()
{
    string expression = "10 + 20 * 30", infix[100], postfix[100] = {};
    int infixCount = -1, postfixCount = 0;

    for (int i = 0; i < expression.length(); i++)
    {
        char newVal = expression[i];
        if (isOperator(newVal) || isNumber(newVal) || isParanthesis(newVal))
        {
            if (infixCount != -1)
            {
                string currVal = infix[infixCount];
                if (infixCount != -1 && (isNumber(currVal[0]) && isNumber(currVal[0])) && (isNumber(newVal) && isNumber(newVal)))
                {
                    infix[infixCount] = (currVal + newVal);
                    continue;
                }
                if (isOperator(newVal) && isOperator(currVal[0]))
                {
                    continue;
                }
            }

            infixCount++;
            infix[infixCount] = newVal;
        }
    }
    infixCount++;
    infix[infixCount] = ")";
    for (int i = 0; i <= infixCount; i++)
    {
        cout << infix[i] << " ";
    }

    for (int i = 0; i <= infixCount; i++)
    {
        char newVal = infix[i][0];
        if (isNumber(newVal))
        {
            postfix[postfixCount++] = infix[i];
        }
        else if (newVal == 40)
        {
            opPush(newVal);
        }
        else if (isOperator(newVal))
        {
            char prevVal = opStack[top];
            if (isOperator(prevVal))
            {
                if (getPriority(prevVal) >= getPriority(newVal))
                {
                    postfix[postfixCount++] = prevVal;
                    opPop();
                }
                opPush(newVal);
            }
            else
            {
                opPush(newVal);
            }
        }
        if (newVal == 41)
        {
            while (true)
            {
                char val = opPop();
                if (val == 40)
                {
                    break;
                }
                postfix[postfixCount++] = val;
            }
        }
    }

    cout << "Postfix " << postfixCount << endl;
    for (int i = 0; i < postfixCount; i++)
    {
        cout << postfix[i] << " ";
    }

    for (int i = 0; i <= postfixCount; i++)
    {
        char newVal = postfix[i][0];
        if (isNumber(newVal))
        {
            ansPush(stoi(postfix[i]));
        }
        else
        {
            int secondVal = ansPop();
            int firstVal = ansPop();
            ansPush(operations(firstVal, secondVal, newVal));
        }
    }
    cout << endl;
    cout << "Result: " << ansStack[0];
}