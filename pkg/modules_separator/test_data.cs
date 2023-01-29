// <Module>
internal class <Module>
{
}

// AccessControl
using System.Collections.Generic;
using KSerialization;
using STRINGS;
using UnityEngine;

[SerializationConfig(MemberSerialization.OptIn)]
[AddComponentMenu("KMonoBehaviour/scripts/AccessControl")]
public class AccessControl : KMonoBehaviour, ISaveLoadable, IGameObjectEffectDescriptor
{
	public enum Permission
	{
		Both,
		GoLeft,
		GoRight,
		Neither
	}

	[MyCmpGet]
	private Operational operational;
}

// YellowAlertMonitor
public class YellowAlertMonitor : GameStateMachine<YellowAlertMonitor, YellowAlertMonitor.Instance>
{
	public new class Instance : GameInstance
	{
		public Instance(IStateMachineTarget master)
			: base(master)
		{
		}

		public void EnableYellowAlert()
		{
		}
	}
}
