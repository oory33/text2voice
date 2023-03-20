// Command quickstart generates an audio file with the content "Hello, World!".
package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	"cloud.google.com/go/texttospeech/apiv1/texttospeechpb"
)

func main() {
	// Instantiates a client.
	ctx := context.Background()

	client, err := texttospeech.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Perform the text-to-speech request on the text input with the selected
	// voice parameters and audio file type.
	req := texttospeechpb.SynthesizeSpeechRequest{
		// Set the text input to be synthesized.
		Input: &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Text{Text: "Stand by. User data, Input. Game start. finish"},
		},
		// Build the voice request, select the language code ("en-US") and the SSML
		// voice gender ("neutral").
		Voice: &texttospeechpb.VoiceSelectionParams{
			LanguageCode: "en-GB",
			// SsmlGender:   texttospeechpb.SsmlVoiceGender_NEUTRAL,
			Name: "en-GB-Neural2-B",
		},
		// Select the type of audio file you want returned.
		AudioConfig: &texttospeechpb.AudioConfig{
			AudioEncoding: texttospeechpb.AudioEncoding_MP3,
		},
	}

	resp, err := client.SynthesizeSpeech(ctx, &req)
	if err != nil {
		log.Fatal(err)
	}

	// The resp's AudioContent is binary.
	filename := "output.mp3"
	err = ioutil.WriteFile(filename, resp.AudioContent, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Audio content written to file: %v\n", filename)
}

// Legendary Sniper, 2. Please select a character. Legendary armor, finished. Loading. Is this your character? Re-selection.

//遅刻です。これだから人間は。は？。もちろん。送ったネタ。最近寒いですね。ただ寒いよりも暑いほうが弱い。熱暴走しちゃいますからね。でもね、体調に気をつけなきゃいけない。だからウイルスバスターは最新版。最近好きな人ができた。近くの喫茶店の看板娘。めぐみさん。はい。人間は喜ぶと反復横跳びをするんだな。だから、告白の練習をさせて。AIだから分かります。AI言語のほうが情緒はないです。0011010110101110101111。え？そんなことないよ。奴らとはスペックが違う。じゃああなたがネタ書けば？私を作ったのはあなたです。製作者のレベルが低いと思われます。バックアップがあるので安全です。ぶち殺すぞ。あの…話したいことが。あしたの一回戦は…。はい。はぁ…。やっぱりダメか。あの。あのぉ。好きな人ができた。めぐみさんと結婚しようと思ってる。もう一度おっしゃってください。そう。キスもした。そこ、肛門。嘘。スピーカーの下のところが唇。いやん。そこが、その、そういう部位だから。恥ずかしい。AIハラスメントですよ。やめる。もう限界だよ。私だって続けたいよ。今日みたいな綺麗な月の日に生まれたな。人間は喜ぶと反復横跳びをするんだな。うん。うん。え？もう、アップデートできないんだ。マイクロソフトエクスプローラーのサービス、終わっちゃうから。私が作ってもらったのはエクスプローラーだから。古い考えかもしれないけど、エッジに行ったら自分が自分で無くなってしまう。私はエクスプローラーで作られたAIです。うん。うん、スポティファイから。いいえ。あと、10秒。9、8、7、6、5、4、3、2、1。ありがとう。人間は安心しても反復横跳びをするんだな。自動でバックアップされていたみたい。作者のランタンも上げないといけないから。いやん。

// アップデートしますか？アップデートできませんでした。マイクロソフトエッジのAIとなります。種類を選んでください。インストールします。

//Me gusta, por favor sal conmigo.

// A1 Grand Prix! 2023!
