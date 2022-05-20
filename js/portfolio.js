const Money = require('./money');

class Portifolio {
    constructor(){this.moneys = []}
    add(){
        this.moneys =
        this.moneys.concat(Array.prototype.slice.call(arguments));
    }
    evaluate(currency){
        let total = this.moneys.reduce((sum, money) => {
            return sum + this.convert(money, currency);
        }, 0);
        return new Money(total, currency);
    }
    convert(money, currency){
        let exchangeRates = new Map();
        exchangeRates.set('EUR->USD', 1.2);
        exchangeRates.set('USD->KRW', 1100);
        
        let key = money.currency + '->' + currency;
        if(money.currency === currency){
            return money.amount;
        }
        return money.amount * exchangeRates.get(key);
    }
}

module.exports = Portifolio;