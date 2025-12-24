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
                <p>シニアプラットフォームエンジニア。</p>
                <p>フリーランスでのWeb受託開発からキャリアをスタートし、スタートアップでの新規事業立ち上げ、受託開発企業での多様な案件対応を経験。</p>
                <p>その後、クラウドファンディング企業にてソフトウェアエンジニアとして通知基盤や認証認可基盤などの基盤システムの設計・開発に従事。</p>
                <p>現在はHRテック企業にてシニアプラットフォームエンジニアとして、プロダクト基盤開発に従事。</p>
                <p><strong><em>「組織・ソフトウェアの価値を長期的に向上させること」</em></strong>を目指して、自らが関わる組織とプロダクトの成長に貢献していきたいと日々思っています。</p>
                <p>システムアーキテクチャに関心があります。</p>
                <p>趣味はOSS開発とジムと観葉植物です。</p>
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
                                <td>GitHub</td>
                                <td><a href="https://github.com/bmf-san">bmf-san</a></td>
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
                                <td>個人ブログ</td>
                                <td><a href="https://bmf-tech.com">bmf-tech.com</a></td>
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
                    <h2>主な経験</h2>
                    <h3>技術・アーキテクチャ</h3>
                    <ul>
                        <li>分散システム・マイクロサービスアーキテクチャ設計</li>
                        <li>基盤システム開発</li>
                        <li>クラウドインフラ構築・運用</li>
                        <li>システムパフォーマンス最適化</li>
                        <li>Go・PHPによる大規模バックエンド開発</li>
                        <li>React・Vue.jsによるフロントエンド開発</li>
                    </ul>
                    <h3>チーム・組織マネジメント</h3>
                    <ul>
                        <li>開発チーム立ち上げ・技術組織マネジメント</li>
                        <li>チームビルディング・プロジェクトマネジメント</li>
                        <li>スクラム開発の導入・運用</li>
                        <li>採用活動・オンボーディング支援</li>
                        <li>技術カルチャー推進・技術戦略策定</li>
                    </ul>
                    <h3>コミュニティ・アウトプット</h3>
                    <ul>
                        <li>OSS開発</li>
                        <li>技術カンファレンス登壇</li>
                        <li>技術書籍執筆（Zenn）</li>
                        <li>個人ブログ運営（bmf-tech.com）</li>
                        <li>キャリア支援・技術メンタリング</li>
                    </ul>
                </div>
                <div class="margin-top-3rem">
                    <h2>職歴</h2>
                    <p><em>株式会社SmartHR/シニアプラットフォームエンジニア</em><br><span class="font-size-small">2024年6月-現在</span></p>
                    <p class="font-size-small">HRテック領域でのプラットフォーム開発・運用、システムアーキテクチャ設計・リアーキテクチャリング、パフォーマンス最適化プロジェクトへの参画。</p>
                    <br>
                    <p><em>株式会社マクアケ/シニアソフトウェアエンジニア・リードエンジニア</em><br><span class="font-size-small">2018年12月-2024年5月</span></p>
                    <p class="font-size-small">システム基盤開発チームのリードエンジニアとしてビジネス成長に伴うシステムスケーラビリティ課題の解決、認証基盤・認可基盤・通知基盤等の基盤システム開発をリード。</p>
                    <br>
                    <p><em>株式会社イノベーター・ジャパン/プロダクトマネージャー兼エンジニア</em><br><span class="font-size-small">2017年5月-2018年9月</span></p>
                    <p class="font-size-small">求人関連の新規サービス立ち上げ（企画・プロトタイプ開発）、CMSとECの機能を持つパッケージ開発、キャンペーンサイトの受託開発。</p>
                    <br>
                    <p><em>スタートアップ企業/Webエンジニア</em><br><span class="font-size-small">2016年5月-2017年3月</span></p>
                    <p class="font-size-small">受託開発企業にて求人サービスの機能開発・UI実装、効果計測によるサービスグロース。</p>
                    <br>
                    <p><em>フリーランス/Web開発エンジニア</em><br><span class="font-size-small">2015年6月-2016年5月</span></p>
                    <p class="font-size-small">Webの受託開発案件に従事、多様なクライアント・プロジェクトでの開発経験。</p>
                </div>
                <div class="margin-top-3rem">
                    <h2>OSS</h2>
                    <p><a href="https://github.com/bmf-san/ggc" target="_blank">ggc</a></p>
                    <p class="font-size-small">Gitクライアントツール</p>
                    <p><a href="https://github.com/bmf-san/gobel-api" target="_blank">gobel-api</a></p>
                    <p class="font-size-small">Go製Headless CMS</p>
                    <p><a href="https://github.com/bmf-san/goblin" target="_blank">goblin</a></p>
                    <p class="font-size-small">Go製HTTP Router</p>
                    <p><a href="https://github.com/bmf-san/gondola" target="_blank">gondola</a></p>
                    <p class="font-size-small">Go製リバースプロキシ</p>
                    <p><a href="https://github.com/bmf-san/go-clean-architecture-web-application-boilerplate" target="_blank">go-clean-architecture-web-application-boilerplate</a></p>
                    <p class="font-size-small">GoでのClean Architecture実装例</p>
                    <p><a href="https://github.com/bmf-san/resume-manager" target="_blank">resume-manager</a></p>
                    <p class="font-size-small">履歴書・職務経歴書管理ツール</p>
                    <p><a href="https://github.com/bmf-san/spec-by-bmf-san" target="_blank">spec-by-bmf-san</a></p>
                    <p class="font-size-small">オリジナルのプラクティスを仕様として管理</p>
                </div>
                <div class="margin-top-3rem">
                    <h2>書籍</h2>
                    <p><a href="https://zenn.dev/books/3f41c5cd34ec3f" target="_blank">net/httpでつくるHTTPルーター自作入門</a><br><span class="font-size-small">2022</span></p>
                    <p class="font-size-small">Goで開発したHTTPルーターの知見を元に、標準パッケージのみでHTTPルーターを実装する流れを本にしました。</p>
                </div>
                <div class="margin-top-3rem">
                    <h2>登壇</h2>
                    <p>全ての登壇資料は<a href="http://speakerdeck.com/bmf_san" target="_blank">Speaker Deck</a>に掲載しています。</p>
                    <p><em><span>アーキテクチャConference 2025</span></em><br><a href="https://speakerdeck.com/bmf_san/wan-bi-woqiu-menaiyi-si-jue-ding-akusesuzhi-yu-ji-pan-niokeruzhi-yue-tonoxiang-kihe-ifang" target="_blank">完璧を求めない意思決定-アクセス制御基盤における制約との向き合い方</a></p>
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
