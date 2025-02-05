{{ define "headline" }}
<div class="headline background-color-main color-text-reverse">
    <div class="container">
        <div class="row">
            <div class="col">
                <h1 class="text-align-center color-text-reverse">プロフィール</h1>
            </div>
        </div>
    </div>
</div>
{{ end }}
{{ define "body" }}
<div class="container-readable margin-top-2rem">
    <div class="row">
        <div class="col margin-top-1rem margin-bottom-1rem">
            <div class="text-align-center">
                <img src="/profile.png" alt="プロフィール画像" style="max-width: 300px;">
            </div>
            <div class="text-align-left">
                <h1 style="margin-top: 0; font-size: xx-large"><strong>Kenta Takeuchi</strong></h1>
                <p>ソフトウェアエンジニア。<br></p>
                <p>2015年に大学卒業後、フリーランス、スタートアップ、受託開発会社等を経て、現在は事業会社の開発部門にてソフトウェアエンジニアとして従事しています。</p>
                <p><strong><em>「組織・ソフトウェアの価値を長期的に向上させること」</em></strong>を目指して、自らが関わる組織とプロダクトの成長に貢献していきたいと日々思っています。</p>
                <p>自分が使いたいアプリケーションのコードを書くのが好きです。<br>趣味はコーディングとジムと観葉植物です。</p>
                <div class="margin-top-3rem">
                    <table>
                        <thead>
                            <tr>
                                <th>Key</th>
                                <th>Value</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr>
                                <td>Email</td>
                                <td>bmf.infomation@gmail.com</td>
                            </tr>
                            <tr>
                                <td>Twitter</td>
                                <td><a href="https://twitter.com/bmf_san">bmf_san</a></td>
                            </tr>
                            <tr>
                                <td>LinkedIn</td>
                                <td><a href="www.linkedin.com/in/bmf-san">bmf-san</a></td>
                            </tr>
                            <tr>
                                <td>GitHub</td>
                                <td><a href="https://github.com/bmf-san">bmf-san</a></td>
                            </tr>
                            <tr>
                                <td>DEV</td>
                                <td><a href="https://dev.to/bmf_san">bmf_san</a></td>
                            </tr>
                            <tr>
                                <td>Zenn</td>
                                <td><a href="https://zenn.dev/bmf_san">bmf_san</a></td>
                            </tr>
                            <tr>
                                <td>Qiita</td>
                                <td><a href="https://qiita.com/bmf_san">bmf_san</a></td>
                            </tr>
                            <tr>
                                <td>Speaker Deck</td>
                                <td><a href="https://speakerdeck.com/bmf_san">bmf_san</a></td>
                            </tr>
                        </tbody>
                    </table>
                </div>
                <div class="margin-top-3rem">
                    <h2>経験</h2>
                    <h3>技術</h3>
                    <ul>
                        <li>GoやPHPを使ったバックエンド開発</li>
                        <li>ReactやVue.jsを使ったフロンドエンド開発</li>
                        <li>分散システムの構築・開発・運用</li>
                        <li>AWSやGCPなどのクラウドベンダーを利用したシステム設計・開発・運用</li>
                    </ul>
                    <h3>チームマネジメント</h3>
                    <ul>
                        <li>チームビルディング</li>
                        <li>チームマネジメント</li>
                        <li>プロジェクトマネジメント</li>
                        <li>スクラム開発の運用</li>
                    </ul>
                    <h3>採用・広報</h3>
                    <ul>
                        <li>採用活動</li>
                        <li>技術カルチャーの推進</li>
                    </ul>
                </div>
                <div class="margin-top-3rem">
                    <h2>職歴</h2>
                    <p><em>株式会社SmartHR/ソフトウェアエンジニア</em><br><span class="font-size-small">2024-現在</span></p>
                    <p class="font-size-small">プロダクト基盤開発に従事。</p>
                    <br>
                    <p><em>株式会社マクアケ/ソフトウェアエンジニア</em><br><span class="font-size-small">2018-2024</span></p>
                    <p class="font-size-small">エンドユーザー・管理者向け機能開発、技術負債解消、システム基盤開発などに従事。</p>
                    <br>
                    <p><em>株式会社イノベーター・ジャパン/ソフトウェアエンジニア</em><br><span class="font-size-small">2017-2018</span></p>
                    <p class="font-size-small">受託・自社開発にてエンドユーザー向け機能開発に従事。</p>
                </div>
                <div class="margin-top-3rem">
                    <h2>資格</h2>
                    <p><span>2023</span> <em>Google Cloud Certified Associate Cloud Engineer</em></p>
                    <p><span>2022</span> <em>Registerd Scrum Master（RSM）</em></p>
                    <p><span>2021</span> <em>AWS Certified Solutions Architect – Associate</em></p>
                    <p><span>2016</span> <em>Webデザイン技能検定2級</em></p>
                </div>
                <div class="margin-top-3rem">
                    <h2>書籍</h2>
                    <p><a href="https://zenn.dev/books/3f41c5cd34ec3f" target="_blank">net/httpでつくるHTTPルーター自作入門</a><br><span class="font-size-small">2022</span></p>
                    <p class="font-size-small">Goで開発したHTTPルーターの知見を元に、標準パッケージのみでHTTPルーターを実装する流れを本にしました。</p>
                </div>
                <div class="margin-top-3rem">
                    <h2>登壇</h2>
                    <p>全ての登壇資料は<a href="http://speakerdeck.com/bmf_san" target="_blank">Speaker Deck</a>に掲載しています。</p>
                    <p><em><span>Go Conference 2024</span></em><br><a href="https://speakerdeck.com/bmf_san/zi-zuo-httprutakaraxin-siiservemuxhe" target="_blank">自作HTTPルーターから新しいServeMuxへ</a></p>
                    <p><em><span>PHPerKaigi 2024</span></em><br><a href="https://speakerdeck.com/bmf_san/gu-kunatutesimatutaphphuremuwakutophpnohasiyonatuhuzhan-lue" target="_blank">古くなってしまったPHPフレームワークとPHPのバージョンアップ戦略</a></p>
                    <p><em><span>Go Conference 2021 Autumn</span></em><br><a href="https://speakerdeck.com/bmf_san/httpdetukuruhttprutazi-zuo-ru-men" target="_blank">net/httpでつくるHTTPルーター自作入門</a></p>
                </div>
                <div class="margin-top-3rem">
                    <h2>ボランティア</h2>
                    <p><em>「ソフトウェアアーキテクチャ・ハードパーツ ―分散アーキテクチャのためのトレードオフ分析」</em>の書籍レビュアとして参加</p>
                    <p><em>「スタッフエンジニアの道 ―優れた技術専門職になるためのガイド」</em>の書籍レビュアとして参加</p>
                    <p><em>「ソフトウェアアーキテクトのための意思決定術　リーダーシップ／技術／プロダクトマネジメントの活用」</em>の書籍レビュアとして参加</p>
                </div>
                <div class="margin-top-3rem">
                    <h2>メディア掲載</h2>
                    <table>
                        <thead>
                            <tr>
                                <th>Key</th>
                                <th>Value</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr>
                                <td><a href="https://freelance.levtech.jp/" target="_blank">レバテックフリーランス</a></td>
                                <td><a href="https://freelance.levtech.jp/guide/detail/1641/" target="_blank">市場価値を高めたいエンジニアに読んでほしい技術ブログ特集</a></td>
                            </tr>
                            <tr>
                                <td><a href="https://dividable.net/" target="_blank">DAINOTE</a></td>
                                <td><a href="https://dividable.net/it-career/engineer/ruby-no-experience" target="_blank">未経験からRubyエンジニアになるには？求人例・転職方法・スキルを紹介</a></td>
                            </tr>
                            <tr>
                                <td><a href="https://assist-all.co.jp/column/" target="_blank">マーケティング情報局</a></td>
                                <td><a href="https://assist-all.co.jp/column/hp/20241104/" target="_blank">プログラマーが提供するホームページ制作！専門性と効果的なWeb戦略を解説 #ホームページ制作プログラマー #ホームページ制作 #プログラマー</a></td>
                            </tr>
                        </tbody>
                    </table>
                </div>
            </div>
        </div>
    </div>
    <div class="row margin-bottom-2rem">
        <div class="col">
                <p>ブログや仕事に関するお問い合わせ等は下記よりお問い合わせください。<br>※返信を約束するものではないことをご了承ください。</p>
                <div>
                    <button style="width:100%" onclick="window.open('https://forms.gle/zZamAmNEtHZGnwZY6')">お問い合わせ</button>
                </div>
        </div>
    </div>
</div>
{{ end }}
