#!/usr/bin/perl
print "Convert your U.S. Dollars to foreign currency\n";
$BritishPounds=0.77;
$Euros=0.92;
$MexicanPesos=18.65;
$Canadiandollar=1.33;
$Solution;
do{
print "How many U.S. Dollars do you want to exchange?";
$dollar=<STDIN>;
chomp($dollar);
print"Enter P (British Pounds), E (Euros), M (Mexican Pesos), or C (Canadian):";
$choice=<STDIN>;
chomp($choice);
if($choice eq "P")
{
    $Solution=$dollar*$BritishPounds;
    printf("Your converted amount is %.2f\n",$Solution);
}
elsif($choice eq "E")
{
    $Solution=$dollar*$Euros;
    printf("Your converted amount is %.2f\n",$Solution);
}
elsif($choice eq "M")
{
    $Solution=$dollar*$MexicanPesos;
    printf("Your converted amount is %.2f\n",$Solution);
}
elsif($choice eq "C")
{
    $Solution=$dollar*$Canadiandollar;
    printf("Your converted amount is %.2f\n",$Solution);
}
printf("Do you want to exchange again (yes/no)?");
$continue=<STDIN>;
chomp($continue);
if($continue=="no")
{
    printf("Thank you for your business");
}
}while($continue ne "no");





