import { Classes, Button, Menu, MenuDivider, MenuItem, ButtonGroup } from "@blueprintjs/core";
import { Popover2 } from "@blueprintjs/popover2";
import React from "react";
import { useTranslation } from "react-i18next";
import { AppThunk, RootState, useAppDispatch, useAppSelector } from "../../Stores/store";
import { userActions } from "../../Stores/userSlice";
import { ImportFromGOODDialog, ImportFromEnkaDialog } from "./Components";
import { viewerActions } from "../../Stores/viewerSlice";
import { Executor, ExecutorSupplier } from "@gcsim/executors";
import ExecutorSettingsButton from "../../Components/Buttons/ExecutorSettingsButton";
import { throttle } from "lodash-es";
import { SimResults } from "@gcsim/types";
import { VIEWER_THROTTLE } from "../Viewer";
import { useHistory } from "react-router";

type Props = {
  exec: ExecutorSupplier<Executor>;
  cfg: string;
  isReady: boolean;
  isValid: boolean;
};

export function runSim(pool: Executor, cfg: string): AppThunk {
  return function (dispatch) {
    console.log("starting run");
    dispatch(viewerActions.start());

    const updateResult = throttle(
      (res: SimResults, hash: string) => {
        dispatch(viewerActions.setResult({ data: res, hash: hash }));
      },
      VIEWER_THROTTLE,
      { leading: true, trailing: true }
    );

    pool.run(cfg, updateResult).catch((err) => {
      dispatch(viewerActions.setError({ error: err }));
    });
  };
}

export const Toolbox = ({ exec, cfg, isReady, isValid }: Props) => {
  const { t } = useTranslation();
  const history = useHistory();

  const [openImport, setOpenGOODImport] = React.useState<boolean>(false);
  const [openImportFromEnka, setOpenImportFromEnka] = React.useState<boolean>(false);
  const { settings } = useAppSelector((state: RootState) => {
    return {
      settings: state.user.data.settings,
    };
  });

  const dispatch = useAppDispatch();
  const toggleTips = () => {
    dispatch(
      userActions.setUserSettings({
        showTips: !settings.showTips,
        showBuilder: settings.showBuilder,
      })
    );
  };

  const run = () => {
    dispatch(runSim(exec(), cfg));
    history.push("/web");
  };

  const toggleBuilder = () => {
    dispatch(
      userActions.setUserSettings({
        showTips: settings.showTips,
        showBuilder: !settings.showBuilder,
      })
    );
  };

  const ToolMenu = (
    <Menu>
      <MenuItem
        icon="help"
        text={settings.showTips ? "Hide Tooltips" : "Show Tooltips"}
        onClick={toggleTips}
      />
      <MenuItem
        icon="people"
        text={settings.showBuilder ? "Hide Builder" : "Show Builder"}
        onClick={toggleBuilder}
      />
      <MenuDivider />
      <MenuItem
          text="Sample Upload"
          icon="helper-management"
          onClick={() => history.push("/sample/upload")}/>
      <MenuItem icon="cut" text="Substat Snippets" disabled />
      <MenuDivider />

      <MenuItem text="Import from GO" icon="import" onClick={() => setOpenGOODImport(true)} />
      <MenuItem text="Import from Enka" icon="import" onClick={() => setOpenImportFromEnka(true)} />
    </Menu>
  );

  return (
    <div className="p-2 wide:ml-2 wide:mr-2 flex flex-row flex-wrap place-items-center gap-x-1 gap-y-1">
      <div className="basis-full wide:basis-0 flex-grow p-1 flex flex-row items-center">
        <ExecutorSettingsButton />
      </div>
      <ButtonGroup className="basis-full wide:basis-2/3 p-1 flex flex-row flex-wrap">
        <Popover2
            content={ToolMenu}
            placement="top"
            className="basis-full md:basis-1/2"
            popoverClassName={Classes.POPOVER_DISMISS}>
          <Button icon="wrench" fill text="Tools" />
        </Popover2>
        <Button
            icon="play"
            intent="primary"
            className="!basis-full md:!basis-1/2"
            onClick={run}
            loading={!isReady}
            disabled={!isValid}
            text={t<string>("simple.run")} />
      </ButtonGroup>
      <ImportFromGOODDialog isOpen={openImport} onClose={() => setOpenGOODImport(false)} />
      <ImportFromEnkaDialog
        isOpen={openImportFromEnka}
        onClose={() => setOpenImportFromEnka(false)}
      />
    </div>
  );
};
