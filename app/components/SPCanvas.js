import { Component, useContext, useEffect } from 'react'
import { initPixi, loader } from '../lib/pixiHelpers'
import { initMatter, stopMatter, registerUpdateCb, initMatterRenderer, unregisterUpdateCb } from '../lib/matterHelpers'
import { Dialog, DialogDetail, Comment, Spacer, moveAdjust, loadRequiredResources } from '../lib/SPGrahic'
import { getDialog, getDialogDetail } from '../lib/api'
import { createMock } from '../lib/createMock'
import Observer from '../lib/observer'
import MainContextProvider from '../contexts/MainContext'
import AuthContextProvider from '../contexts/AuthContext'
import { MainContext } from '../contexts/MainContext'
import { withToast } from '../lib/withToast'

/**
 * このアプリケーションのキャンバス描画
 * @author Ritsuki KOKUBO
 */
class SPCanvas extends Component {
  dialogs
  dialogDetail
  centerSpacer
  comments
  pixi
  matter
  matterRender
  mock
  currentViewMode
  context

  constructor(props) {
    super(props)
    this.dialogs = []
    this.dialogDetail = null
    this.centerSpacer = null
    this.comments = []
    this.pixi = null
    this.matter = null
    this.matterRender = null
    this.mock = null
    this.currentViewMode = 'empty'
    this.context = {}
  }

  componentDidMount() {
    // console.log('_indexs')
    const self = this
    if (process.env.NEXT_PUBLIC_ENV === 'MOCK') {
      this.mock = createMock()
    }

    if (typeof window === 'undefined') {
      return
    }
    this.pixi = initPixi(document.getElementById('spMainCanvas'))
    // console.log('init spcanvas!')
    this.matter = initMatter()
    //デバッグ用レンダラー
    // this.matterRender = initMatterRenderer(document.getElementById('spDebugCanvas'), this.matter.engine)

    registerUpdateCb(self.matter.engine, [
      () => {
        if (self.dialogs.length) {
          moveAdjust(self.pixi, self.matter, self.dialogs)
        }
        if (self.dialogDetail) {
          // moveVectorAdjust([self.dialogDetail])
        }
        if (self.comments) {
          moveAdjust(self.pixi, self.matter, self.comments)
        }
      }
    ])
    // TODO: 条件分岐して詳細viewもつくる
    // listDialogはobserverのgenreのuseeffectでやってくれるからパスでOK
    // this.changeView('listDialog')
  }

  componentWillUnmount() {
    if (process.env.NEXT_PUBLIC_ENV === 'MOCK') {
      this.mock.shutdown()
    }
    unregisterUpdateCb()
    for (let i = 0; i < this.dialogs.length; i++) {
      this.dialogs[i].normalRemoveRender(this.pixi, this.matter.engine.world)
      this.dialogs[i] = null
    }
    this.dialogs = []
    this.pixi = null
    // console.log('unmounted spcanvas')
    this.matter = null
    this.matterRender = null
    this.mock = null
  }

  /**
   * 更新するやつ
   * @param {*} viewMode - one of listDialog, detailDialog
   */
  async changeView(viewMode, id, dialog, genre = 'all', offset = 0, limit = 20) {
    // console.log('changeView')
    // console.log(viewMode)
    const self = this
    const CENTER_X = window.innerWidth / 2
    const CENTER_Y = window.innerHeight / 2
    this.currentViewMode = viewMode
    switch (viewMode) {
      case 'listDialog':
        await loadRequiredResources()
        // console.log('updating!')
        // リセット -=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=
        if (this.dialogs.length > 0) {
          this.dialogs.forEach((dialog) => {
            // dialog.mountModel(this.matter.engine.world)
            // dialog.updateOption({
            //   movement: {
            //     mode: 'Center'
            //   }
            // })
            dialog.normalRemoveRender(this.pixi, this.matter.engine.world)
          })
          this.dialogs = []

          if (this.dialogDetail) {
            this.dialogDetail.normalRemoveRender(this.pixi, this.matter.engine.world)
            this.dialogDetail = null
          }

          if (this.centerSpacer) {
            this.centerSpacer.normalRemoveRender(this.pixi, this.matter.engine.world)
            this.centerSpacer = null
          }

          if (this.comments.length) {
            this.comments.forEach((comment) => {
              comment.normalRemoveRender(this.pixi, this.matter.engine.world)
            })
            this.comments = []
          }

          // return
        }
        // -=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=

        const res = await getDialog({
          genre: genre,
          offset: offset,
          limit: limit
        })
        if (!res) {
          this.props.addToast('サーバーにアクセスできませんでした。', { appearance: 'error' })
          return
        }
        res.schema.forEach((s, i) => {
          const dialog = new Dialog(Math.random() * CENTER_X * 2, Math.random() * CENTER_Y * 2, {
            dialog: s.content
          })
          dialog.presentation.on('pointerdown', () => {
            // self.context = {
            // this.context = {
            //   targetDialog: dialog,
            //   targetDialogData: s
            // }
            // console.log(self.context)
            // do routing here
            this.props.setMode('detail')
            this.props.setDialogID(s.id)
            this.props.setDialog(s)
            location.hash = `dialog/${s.id}`
            // window.setTimeout(() => {
            self.changeView('detailDialog', s.id, dialog)
            // }, 1000)
          })
          window.setTimeout(() => {
            dialog.easingInitRender(this.pixi, this.matter.engine.world)
          }, 100 * i)
          self.dialogs.push(dialog)
        })
        break
      case 'detailDialog':
        // console.log(this.context)
        // if (typeof self.context === 'undefined') { // undefined
        //   return
        // }
        await loadRequiredResources()
        // idだけでfetchしたい
        // const { targetDialog, targetDialogData } = context
        // if (typeof targetDialog === 'undefined' && typeof targetDialogData === 'undefined') {
        //   // console.log('bye')
        //   return
        // }

        /*----------------.
        | 追加データフェッチ |
        `-----------------*/
        const detailRes = await getDialogDetail(id, {
          // genre: 'all',
          limit: 20,
          offset: 0
        })
        if (!detailRes) {
          this.props.addToast('サーバーにアクセスできませんでした。', { appearance: 'error' })
          return
        }
        const dialogData = detailRes.dialog
        const commentDatas = detailRes.comments
        // const targetDialog = new Dialog(Math.random() * 0.8 * CENTER_X * 1.1, Math.random() * 0.8 * CENTER_Y * 1.1, {
        //   dialog: detailRes.content
        // })
        const targetDialog = dialog

        /*--------.
        | 描画実行 |
        `--------*/
        // セリフ詳細表示オブジェクト
        const makeDialogDetail = () => {
          if (self.dialogDetail) {
            self.dialogDetail.normalRemoveRender(self.pixi, self.matter.engine.world)
            self.dialogDetail = null
          }
          self.dialogDetail = new DialogDetail(
            0,
            300,
            {
              // author: targetDialogData.author,
              // title: targetDialogData.title,
              // source: targetDialogData.link,
              // cite: targetDialogData.source
              author: dialogData.author,
              title: dialogData.title,
              source: dialogData.link,
              cite: dialogData.source
            },
            {
              movement: {
                mode: 'None'
              }
            }
          )
          self.dialogDetail.x = CENTER_X + 5
          self.dialogDetail.y =
            CENTER_Y + (targetDialog.height + self.dialogDetail.height) / 2 - self.dialogDetail.height / 2
          self.dialogDetail.easingInitRender(self.pixi, self.matter.engine.world).then(() => {
            self.dialogDetail.unmountModel(self.matter.engine.world)
          })
        }

        // 中央スペーサー
        const makeCenterSpacer = (width, height) => {
          if (self.centerSpacer) {
            return
          }
          self.centerSpacer = new Spacer(
            CENTER_X,
            CENTER_Y,
            {},
            {
              width,
              height
            }
          )
          self.centerSpacer.normalInitRender(this.pixi, this.matter.engine.world)
        }

        const makeComments = () => {
          if (self.comments && self.comments.length > 0) {
            self.comments.forEach((comment) => {
              comment.normalRemoveRender(this.pixi, this.matter.engine.world)
            })
            self.comments = []
          }
          commentDatas.forEach((commentData) => {
            self.comments.push(
              new Comment(Math.random() * CENTER_X * 2, Math.random() * CENTER_Y * 2, {
                comment: commentData.content,
                userName: commentData.user['display_name'],
                userPhoto: '',
                time: '●分前'
              })
            )
          })
          self.comments.forEach((comment) => {
            comment.easingInitRender(this.pixi, this.matter.engine.world)
          })
        }
        self.dialogs.forEach((dialog) => {
          if (dialog == targetDialog) {
            dialog.updateOption({
              movement: {
                mode: 'CenterFix',
                context: {
                  offsetX: 0,
                  offsetY: 0,
                  callback: () => {
                    makeDialogDetail()
                    makeCenterSpacer(
                      Math.max(targetDialog.width, self.dialogDetail.width),
                      targetDialog.height + self.dialogDetail.height
                    )
                    dialog.easingMoveRender(
                      window.innerWidth / 2 -
                        Math.max(targetDialog.width, self.dialogDetail.width) / 2 +
                        targetDialog.width / 2 -
                        5,
                      window.innerHeight / 2 -
                        (targetDialog.height + self.dialogDetail.height) / 2 +
                        targetDialog.height / 2
                    )
                    makeComments()
                  }
                }
              }
            })
          } else {
            dialog.updateOption({
              movement: {
                mode: 'OutOfRange'
              }
            })
          }
        })
        break
      case 'new':
        // console.log('about to new post')
        self.dialogs.forEach((dialog) => {
          dialog.updateOption({
            movement: {
              mode: 'OutOfRange'
            }
          })
        })
        break
      default:
        break
    }
  }

  /**
   * 詳細から元のリストに戻す
   */
  async back() {
    switch (this.currentViewMode) {
      case 'detailDialog':
        this.changeView('listDialog')
        break
      default:
        break
    }
  }

  renderPostedComment(txt) {
    if (txt.length < 1) {
      return
    }
    const CENTER_X = window.innerWidth / 2
    const CENTER_Y = window.innerHeight / 2
    console.log(`spcanvas ${txt}`)
    const comment = new Comment(Math.random() * CENTER_X * 2, Math.random() * CENTER_Y * 2, {
      comment: txt,
      userName: 'あなた',
      userPhoto: '',
      time: '●分前'
    })
    comment.easingInitRender(this.pixi, this.matter.engine.world)
    this.comments.push(comment)
  }

  render() {
    return (
      <>
        <MainContextProvider>
          <Observer
            cb={() => {}}
            changeView={this.changeView}
            shouldUpdate={this.props.shouldUpdate}
            setShouldUpdate={this.props.setShouldUpdate}
            self={this}
            selectedGenre={this.props.selectedGenre}
            mode={this.props.mode}
            setSelectedGenre={this.props.setSelectedGenre}
            cameBack={this.props.cameBack}
            setCameBack={this.props.setCameBack}
            mounted={this.props.mounted}
            setMounted={this.props.setMounted}
            postedCommet={this.props.postedCommet}
            renderPostedComment={this.renderPostedComment}
          />
        </MainContextProvider>
        <div
          id="spDebugCanvas"
          style={{
            position: 'absolute',
            top: 0,
            left: 0,
            height: '100vh',
            width: '100vw',
            zIndex: 0
          }}></div>
        <canvas
          id="spMainCanvas"
          style={{
            position: 'absolute',
            top: 0,
            left: 0,
            height: '100vh',
            width: '100vw',
            zIndex: 0 // TODO: 吟味の余地あり
          }}
        />
      </>
    )
  }
}

export default withToast(SPCanvas)
