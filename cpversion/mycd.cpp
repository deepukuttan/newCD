#include <iostream>
#include <vector>

using namespace std;

vector<string> split(string st,char s){
    int a=0;
    vector<string> out;
    string temp;
    for(int i=0;i<st.length();i++){
        if(s[i]==s || i==st.length()){
            out.insert(out.end(),temp);
            temp="";
        }else{
            temp+=st[i];
        }
    }

    return out;
}

bool isin(vector<string> b,string s){
    bool sts = 0;
    for(auto i=b.begin();i<b.end();i++){
        if(*i==s){
            sts=1;
            break;
        }
    }
    return sts;
}

bool sub(string s,string sub){
    string temp;
    bool sts = 0;
    for(int i=0;i<s.length();i+=sub.length()){
        for(int k=i;k<(i+sun.length());k++){
            temp+=s[k];
            if(temp == sub){
                sts = 1;
                break;
            }
        }
        if(sts){
            break;
        }
    }
    return sts;
}

string dirCheck(vector<string> p,vector<string> c){
    string out;
    vector<string> temp;
    bool err = 0;
    temp = p;
    for(auto i=c.begin();i<c.end();i++){
        if(sub(*i,"..")&& *i != ".."){
            cout << "ERROR: Invalid dir name" << endl;
            err = 1;
            break;
        }
        if(*i=="") continue;
        if(!isin(p,*i) && *i != ".."){
            temp.insert(te,p.end(),*i);
        }else if(!isin(p,*i) && *i == ".."){
            temp.erase(temp.end());
        }
    }

    if(err){
        return "";
    }

    for(auto i=temp.begin();i<temp.end();i++){
        out += *i + "/";
    }

    return out;
}

int main(int argc,char **argv){
    vector<string> cur = split(argv[1].'/');
    vector<string> news = split(argv[2].'/');
    cout << dirCheck(cur,news) << endl;
}