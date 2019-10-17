
if (typeof window.Vue === 'undefined') {
    throw new Error("Vue not fount");
}

// 洗牌算法
function shuffle(array) {
    var copiedArray = array.slice();

    for (var index = copiedArray.length - 1; index >= 0; index--) {
        var randomIndex = Math.floor(Math.random() * (index + 1));
        var randomItem = copiedArray[randomIndex];

        copiedArray[randomIndex] = copiedArray[index];
        copiedArray[index] = randomItem;
    }

    return copiedArray;
}

// 获取词语文字
function getLemmaWords(lemmas) {
    var lemmaWords = [];

    lemmas.forEach(function(lemma) {
        var lemmaWordItems = lemma.name.split('').map(name => ({ name, removed: false, id: Math.random().toString() }));
        lemmaWords.push(...lemmaWordItems);
    });

    return shuffle(lemmaWords);
}

var { lemmas } = window.DATA.data;

var lemmaGame = new Vue({
    el: '#root',
    data: {
        lemmas,
        currentLemma: [],
        randomLemmaWords: getLemmaWords(lemmas),
    },
    computed: {
        layoutRandomLemmaWords: function() {
            var screenWidth = window.innerWidth - 20;
            var layoutRandomLemmaWords = this.randomLemmaWords.slice();

            if (layoutRandomLemmaWords.length % Math.floor(screenWidth / 36) > 0) {
                var needAddBlankLemmaWordNumber = Math.floor(screenWidth / 36) - (layoutRandomLemmaWords.length % Math.floor(screenWidth / 36));
                var blankLemmaWords = new Array(needAddBlankLemmaWordNumber).fill('').map(function(name) {
                    return { name, id: `$_${Math.random()}`, removed: true }
                });

                layoutRandomLemmaWords.push(...blankLemmaWords);
            }

            return layoutRandomLemmaWords;
        }
    },
    mount: function() {

    },
    methods: {
        handleLemmaWordClick: function(lemmaWord) {
            if (/^\$/.test(lemmaWord.id)) {
                return;
            }

            var currentLemma = this.currentLemma.indexOf(lemmaWord.id) > -1
                ? this.currentLemma.filter(item => item !== lemmaWord.id)
                : this.currentLemma.concat([lemmaWord.id]);

            if (currentLemma.length === 4) {
                this.handleValidateCurrentLemma(currentLemma);
                return;
            }

            this.currentLemma = currentLemma;
        },
        handleValidateCurrentLemma(currentLemma) {
            var currentLemmaText = currentLemma.map(id => this.randomLemmaWords.find(item => item.id === id).name).join('');

            if (this.lemmas.map(lemma => lemma.name).indexOf(currentLemmaText) === -1) {
                alert('您选择的成语有错误');
                return;
            }

            var newLemmas = this.lemmas.map((lemma, index) => {
                return Object.assign({}, lemma, {
                    removed: lemma.removed || (this.lemmas.map(lemma => lemma.name).indexOf(currentLemmaText) === index ? true : false)
                });
            });

            var newRandomLemmaWords = this.randomLemmaWords.map(lemmaWord => {
                return Object.assign({}, lemmaWord, {
                    removed: lemmaWord.removed || currentLemma.indexOf(lemmaWord.id) > -1 ? true : false
                });
            });

            this.currentLemma = [];
            this.lemmas = newLemmas;
            this.randomLemmaWords = newRandomLemmaWords;
        },
        handleLemmaTip() {
            var ableLemmas = this.lemmas.filter(lemma => !lemma.removed);

            if (ableLemmas.length === 0) {
                return;
            }

            let randomTipLemmaWords = ableLemmas[Math.floor(Math.random() * ableLemmas.length)].name.split('');

            var newRandomLemmaWords = this.randomLemmaWords.map(word => {
                if (!word.removed && randomTipLemmaWords.indexOf(word.name) > -1) {
                    randomTipLemmaWords.splice(randomTipLemmaWords.indexOf(word.name), 1);
                    return Object.assign({}, word, { isTip: true });
                }
                return word;
            });

            this.randomLemmaWords = newRandomLemmaWords;

            window.setTimeout(() => {
                this.randomLemmaWords = this.randomLemmaWords.map(word => Object.assign({}, word, { isTip: false }))
            }, 2000);
        },
        handleLemmaAgain() {
            window.location.reload();
        }
    }
});


